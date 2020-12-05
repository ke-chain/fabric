ALPINE_VER ?= 3.12

BUILD_DIR ?= build


EXTRA_VERSION ?= $(shell git rev-parse --short HEAD)
PROJECT_VERSION=$(BASE_VERSION)-snapshot-$(EXTRA_VERSION)

RELEASE_EXES = orderer 

# defined in common/metadata/metadata.go
METADATA_VAR = Version=$(BASE_VERSION)
METADATA_VAR += CommitSHA=$(EXTRA_VERSION)
METADATA_VAR += BaseDockerLabel=$(BASE_DOCKER_LABEL)
METADATA_VAR += DockerNamespace=$(DOCKER_NS)

GO_VER = 1.14.12
GO_TAGS ?=

RELEASE_IMAGES = orderer

PKGNAME = github.com/ke-chain/fabric

pkgmap.orderer        := $(PKGNAME)/cmd/orderer


.DEFAULT_GOAL := docker

include docker-env.mk
include gotools.mk


.PHONY: docker
docker: clean $(RELEASE_IMAGES:%=%-docker)

.PHONY: $(RELEASE_IMAGES:%=%-docker)
$(RELEASE_IMAGES:%=%-docker): %-docker: $(BUILD_DIR)/images/%/$(DUMMY)

$(BUILD_DIR)/images/orderer/$(DUMMY): BUILD_ARGS=--build-arg GO_TAGS=${GO_TAGS}

$(BUILD_DIR)/images/%/$(DUMMY):
	@echo "Building Docker image $(DOCKER_NS)/fabric-$*"
	@mkdir -p $(@D)
	$(DBUILD) -f images/$*/Dockerfile \
		--build-arg GO_VER=$(GO_VER) \
		--build-arg ALPINE_VER=$(ALPINE_VER) \
		$(BUILD_ARGS) \
		-t $(DOCKER_NS)/fabric-$* ./$(BUILD_CONTEXT)
	@touch $@



.PHONY: clean
clean: 
	-@rm -rf $(BUILD_DIR)

.PHONY: $(RELEASE_EXES)
$(RELEASE_EXES): %:  clean $(BUILD_DIR)/bin/%
	

$(BUILD_DIR)/bin/%: GO_LDFLAGS = $(METADATA_VAR:%=-X $(PKGNAME)/common/metadata.%)
$(BUILD_DIR)/bin/%:
	@echo "Building $@"
	@mkdir -p $(@D)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOBIN=$(abspath $(@D)) go build -o ./build/bin/orderer -tags "$(GO_TAGS)" -ldflags "$(GO_LDFLAGS)" $(pkgmap.$(@F))
	@touch $@
