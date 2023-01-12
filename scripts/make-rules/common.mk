SHELL := /bin/bash
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

ifeq ($(origin ROOT_DIR), undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../../ && pwd -P))
endif

ifeq ($(origin OUTPUT_DIR), undefined)
OUTPUT_DIR := $(ROOT_DIR)/_output
endif

ifeq ($(origin TOOLS_DIR), undefined)
TOOLS_DIR := $(OUTPUT_DIR)/tools
$(shell mkdir -p $(TOOLS_DIR))
endif

ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif

# Set the git version
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

# check if the git tree is dirty, default is dirty
GIT_TREE_STATE := dirty
ifeq (, $(shell git status --porcelain 2>/dev/null))
GIT_TREE_STATE := clean
endif

# Get commit id
GIT_COMMIT := $(shell git rev-parse HEAD)

# Minimum test coverage
ifeq ($(origin COVERAGE),undefined)
COVERAGE := 60
endif

PLATFORMS ?= linux_amd64 linux_arm64

