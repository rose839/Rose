# Build all by default, even if it's not first
.DEFAULT_GOAL = all

.PHONY: all
all: build

# ==============================================================================
# Includes
include scripts/make-rules/golang.mk

# ==============================================================================
# Build options

ROOT_PACKAGE=github.com/rose839/Rose
VERSION_PACKAGE=github.com/rose839/Rose/pkg/version

# ==============================================================================
# Usage
define USAGE_OPTIONS

Options:
  DEBUG            Whether to generate debug symbols. Default is 0.
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build/build.multiarch
                   Example: make build BINS="iam-apiserver iam-authz-server"
  IMAGES           Backend images to make. Default is all of cmd starting with iam-.
                   This option is available when using: make image/image.multiarch/push/push.multiarch
                   Example: make image.multiarch IMAGES="iam-apiserver iam-authz-server"
  REGISTRY_PREFIX  Docker registry prefix. Default is marmotedu. 
                   Example: make push REGISTRY_PREFIX=ccr.ccs.tencentyun.com/marmotedu VERSION=v1.6.2
  PLATFORMS        The multiple platforms to build. Default is linux_amd64 and linux_arm64.
                   This option is available when using: make build.multiarch/image.multiarch/push.multiarch
                   Example: make image.multiarch IMAGES="iam-apiserver iam-pump" PLATFORMS="linux_amd64 linux_arm64"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

# ==============================================================================
# Targets

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "Usage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n s/^##//p $< | column -t -s ":" | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
