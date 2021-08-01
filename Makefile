VERSION := $(shell cat VERSION)
REGISTRY := registry.gitlab.com/frenchtoasters/vrliingest
IMAGE_REPO := $(REGISTRY)/familytree
IMAGE_TAG := $(VERSION)
BUILD_DIR := build

.DEFAULT_GOAL := build-local

.PHONY: revendor
revendor:
	@env GO111MODULE=on go mod vendor -v
	@env GO111MODULE=on go mod tidy -v

.PHONY: update-dependencies
update-dependencies:
	@env GO111MODULE=on go get -u
	@make revendor

.PHONY: build-local
build-local:
	@env LOCAL_BUILD=1 .ci/build

.PHONY: docker-image
docker-image:
	@if [[ ! -f $(BIN_DIR)/linux-amd64/etcdbrctl ]]; then echo "No binary found. Please run 'make build'"; false; fi
	@docker build -t $(IMAGE_REPOSITORY):$(IMAGE_TAG) -f $(BUILD_DIR)/Dockerfile --rm .

.PHONY: docker-push
docker-push:
	@if ! docker images $(IMAGE_REPOSITORY) | awk '{ print $$2 }' | grep -q -F $(IMAGE_TAG); then echo "$(IMAGE_REPOSITORY) version $(IMAGE_TAG) is not yet built. Please run 'make docker-image'"; false; fi
	@docker push $(IMAGE_REPOSITORY):$(IMAGE_TAG)

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)/