# The development version of clang is distributed as the 'clang' binary,
# while stable/released versions have a version number attached.
# Pin the default clang to a stable version.
CLANG ?= clang-11
CFLAGS := -target bpf -O2 -g -Wall -Werror $(CFLAGS)

# Obtain an absolute path to the directory of the Makefile.
# Assume the Makefile is in the root of the repository.
REPODIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
UIDGID := $(shell stat -c '%u:%g' ${REPODIR})

REPO := $(shell cat ${REPODIR}/testdata/docker/REPO)
VERSION := $(shell cat ${REPODIR}/testdata/docker/VERSION)

# clang <8 doesn't tag relocs properly (STT_NOTYPE)
# clang 9 is the first version emitting BTF
TARGETS := \
	testdata/loader-clang-7 \
	testdata/loader-clang-9 \
	testdata/loader-clang-11 \
	testdata/invalid_map \
	testdata/raw_tracepoint \
	testdata/invalid_map_static \
	testdata/initialized_btf_map \
	testdata/strings \
	internal/btf/testdata/relocs

.PHONY: all clean docker-all docker-shell

.DEFAULT_TARGET = docker-all

# Build all ELF binaries using a Dockerized LLVM toolchain.
docker-all:
	docker run --rm --user "${UIDGID}" \
		-v "${REPODIR}":/ebpf -w /ebpf \
		"${REPO}:${VERSION}" \
		make all

# (debug) Drop the user into a shell inside the Docker container as root.
docker-shell:
	docker run --rm -ti \
		-v "${REPODIR}":/ebpf -w /ebpf \
		"${REPO}:${VERSION}"

clean:
	-$(RM) testdata/*.elf
	-$(RM) internal/btf/testdata/*.elf

all: $(addsuffix -el.elf,$(TARGETS)) $(addsuffix -eb.elf,$(TARGETS))

testdata/loader-%-el.elf: testdata/loader.c
	$* $(CFLAGS) -mlittle-endian -c $< -o $@

testdata/loader-%-eb.elf: testdata/loader.c
	$* $(CFLAGS) -mbig-endian -c $< -o $@

%-el.elf: %.c
	$(CLANG) $(CFLAGS) -mlittle-endian -c $< -o $@

%-eb.elf : %.c
	$(CLANG) $(CFLAGS) -mbig-endian -c $< -o $@

# Usage: make KDIR=/path/to/foo vmlinux-btf.gz
vmlinux-btf.gz: $(KDIR)/vmlinux
	objcopy --dump-section .BTF=/dev/stdout "$<" | gzip > "$@"
