# Makefile for jquery-proxy
# Examples:
#   make build                       # local dev build -> ./build/jquery-proxy
#   make docker-build                # builds image from local artifact
#   make docker-buildx               # builds multi-arch image (multi-stage)
#   make docker-push                 # push local-arch image
#   make docker-pushx                # push multi-arch image
#   make run                         # run locally
#   make clean

SHELL := bash
.ONESHELL:

# ---- Project metadata ----
BIN := jquery-proxy
MODULE_PATH := github.com/bnassif/$(BIN)
VERSION ?= $(shell (git describe --tags --abbrev=0 2>/dev/null) || echo 1.0-1)

# ---- Project ----
ROOT_DIR     := $(abspath .)
OUT_DIR      := $(ROOT_DIR)/dist
BIN_OUT      := $(OUT_DIR)/$(BIN)

# ---- Go build flags ----
export CGO_ENABLED ?= 0
GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
LDFLAGS := -X $(MODULE_PATH)/app.Version=$(VERSION)

# ---- Image ----
IMAGE ?= ghcr.io/bnassif/jquery-proxy
TAG   ?= $(VERSION)
PLATFORMS ?= linux/amd64,linux/arm64
DOCKER_FILE := Dockerfile
DOCKER_EXPORT := "$(OUT_DIR)/$(BIN)_$(VERSION).tar"

DOCKER_FLAGS := -t $(IMAGE):$(TAG) -f $(DOCKER_FILE) --label "Version=$(VERSION)"
DOCKERX_FLAGS := $(BUILD_FLAGS) --platform $(PLATFORMS)


.DEFAULT_GOAL := help

# ---- Helpers ----
define need
	@command -v $(1) >/dev/null 2>&1 || { echo "ERROR: missing dependency: $(1)"; exit 1; }
endef

# ---- Targets ----
.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS=":.*##"; print "Usage: make <target>\n"} /^[a-zA-Z0-9_.-]+:.*##/ { printf "  %-18s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: bin
bin: ## Build local binary -> ./build/jquery-proxy
	$(call need,go)
	mkdir -p "$(OUT_DIR)"
	go build -trimpath -ldflags '$(LDFLAGS)' -o "$(BIN_OUT)" ./

.PHONY: release
release: bin docker docker_latest docker_push

.PHONY: docker
docker: ## Build image using artifact-copy Dockerfile
	$(call need,docker)
	echo "Building with flags: $(DOCKER_FLAGS)"
	docker build $(DOCKER_FLAGS) -f $(DOCKER_FILE) .

.PHONY: docker_latest
docker_latest: ## Tag the Docker Image as Latest
	$(call need,docker)
	docker image tag $(IMAGE):$(TAG) $(IMAGE):latest

.PHONY: docker_push
docker_push: ## Push the single-arch image
	$(call need,docker)
	docker push $(IMAGE):$(TAG)

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf "$(OUT_DIR)"
