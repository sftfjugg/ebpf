#!/bin/bash
# Test the current package under a different kernel.
# Requires virtme and qemu to be installed.

set -eu
set -o pipefail

if [[ "${1:-}" = "--in-vm" ]]; then
  shift

  mount -t bpf bpf /sys/fs/bpf
  export CGO_ENABLED=0
  export GOFLAGS=-mod=readonly
  export GOPATH=/run/go-path
  export GOPROXY=file:///run/go-root/pkg/mod/cache/download
  export GOCACHE=/run/go-cache

  goBinary=/usr/local/bin/go
  if [[ $# == 2 ]]; then
    goBinary=$GOPATH/bin/$1
    HOME=/run
    shift
  fi

  echo Running tests...
  $goBinary test -coverprofile="$1/coverage.txt" -covermode=atomic -v ./...
  touch "$1/success"
  exit 0
fi

# Pull all dependencies, so that we can run tests without the
# vm having network access.
go mod download

# Use sudo if /dev/kvm isn't accessible by the current user.
sudo=""
if [[ ! -r /dev/kvm || ! -w /dev/kvm ]]; then
  sudo="sudo"
fi
readonly sudo

readonly kernel_version="${1:-}"
if [[ -z "${kernel_version}" ]]; then
  echo "Expecting kernel version as first argument"
  exit 1
fi

# get current go stable versions
versions=$(curl https://golang.org/dl/?mode=json | jq '.[] | select(.stable == true) | .version')
for version in $versions; do
  goVersion="${version//\"}"
  go install golang.org/dl/$goVersion
  $goVersion download
done

readonly kernel="linux-${kernel_version}.bz"
readonly output="$(mktemp -d)"
readonly tmp_dir="$(mktemp -d)"

test -e "${tmp_dir}/${kernel}" || {
  echo Fetching ${kernel}
  curl --fail -L "https://github.com/newtools/ci-kernels/blob/master/${kernel}?raw=true" -o "${tmp_dir}/${kernel}"
}

for version in $versions; do
  goVersion="${version//\"}"
  echo Testing on ${kernel_version} with go version $goVersion
  $sudo virtme-run --kimg "${tmp_dir}/${kernel}" --memory 512M --pwd \
    --rwdir=/run/output="${output}" \
    --rodir=/run/go-path="$(go env GOPATH)" \
    --rwdir=/run/go-cache="$(go env GOCACHE)" \
    --rwdir=/run/sdk="${HOME}/sdk" \
    --script-sh "$(realpath "$0") --in-vm ${goVersion} /run/output"

  if [[ ! -e "${output}/success" ]]; then
    echo "Test failed on ${kernel_version} with go version ${goVersion}"
    exit 1
  else
    echo "Test successful on ${kernel_version} with go version ${goVersion}"
    if [[ -v CODECOV_TOKEN ]]; then
      curl --fail -s https://codecov.io/bash > "${tmp_dir}/codecov.sh"
      chmod +x "${tmp_dir}/codecov.sh"
      "${tmp_dir}/codecov.sh" -f "${output}/coverage.txt"
    fi
  fi

done

$sudo rm -r "${output}"
