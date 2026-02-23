# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.DEFAULT_GOAL := help
.SUFFIXES: # remove legacy builtin suffixes to allow easier make debugging
SHELL = /usr/bin/env bash

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = " *:.*## *"; printf "\nUsage:\n  make \033[36m<target>\033[0m [\033[32mVAR\033[0m=val]\n"} /^##.*:##/ { gsub(/^## /,"",$$1) ; printf "  \033[32m%s\033[0m (\"%s\")\n    └ %s\n", $$1, ENVIRON[$$1], $$2} /^[.a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# If GOARCH is not set in the env, find it
GOARCH ?= $(shell go env GOARCH)

#
# ==== ARGS =====
#

##@ Environment args

## DOCKER :## Container build tool compatible with `docker` API
DOCKER ?= docker

## KUBECTL :## Tool compatible with `kubectl` API
KUBECTL ?= kubectl

## PLATFORM :## Platform for builds
PLATFORM ?= linux/$(GOARCH)

## BUILD_ARGS :## Additional args for builds
BUILD_ARGS ?=

## CONTROLLER_TAG :## Image tag for controller image build and deploy
CONTROLLER_TAG ?= cosi-controller:latest

## SIDECAR_TAG :## Image tag for sidecar image build
SIDECAR_TAG ?= cosi-provisioner-sidecar:latest

export

##@ Core (Basic)

.PHONY: all
.NOTPARALLEL: all # all generators must run before build
all: prebuild build ## Build all container images, plus their prerequisites (faster with 'make -j')

.PHONY: lint
lint: kubeapi-lint eof-newline-lint dockerfiles-lint shell-lint ## Run all linters (suggest `make -k`)
golangci-lint:
	$(GOLANGCI_LINT) run $(GOLANGCI_LINT_RUN_OPTS) --config $(CURDIR)/.golangci.yaml
kubeapi-lint: kube-api-linter
	cd client/apis && $(KUBEAPI_LINT) run --config $(CURDIR)/client/.kubeapilint.yaml
dockerfiles-lint:
	hack/tools/lint-dockerfiles.sh $(HADOLINT_VERSION)
eof-newline-lint:
	hack/lint-eof-newline.sh

.PHONY: shell-lint
shell-lint: shellcheck
	$(SHELLCHECK) $(shell git ls-files -- '*.sh' ':(exclude)vendor/*')

.PHONY: lint-fix
lint-fix: golangci-lint-fix ## Run all linters and perform fixes where possible (suggest `make -k`)
golangci-lint-fix:
	$(GOLANGCI_LINT) run $(GOLANGCI_LINT_RUN_OPTS) --config $(CURDIR)/.golangci.yaml --fix

.PHONY: test
test: .test.proto .test.go ## Run all unit tests including fmt
.test.go: fmt
	go test -v -cover ./...
.PHONY: .test.proto
.test.proto: # gRPC proto has a special unit test
	$(MAKE) -C proto check

.PHONY: clean
clean: ## Clean build environment
	$(MAKE) -C proto clean

.PHONY: clobber
clobber: ## Clean build environment and cached tools
	$(MAKE) -C proto clobber
	rm -rf $(TOOLBIN)
	rm -rf $(CURDIR)/.cache

##@ Development (Advanced)

.PHONY: prebuild .gen .doc-vendor
.gen: generate codegen # can be done in parallel with 'make -j'
.doc-vendor: docs vendor # can be done in parallel
.NOTPARALLEL: prebuild # codegen must be finished before fmt
prebuild: .gen fmt .doc-vendor ## Run all pre-build prerequisite steps (faster with 'make -j')

.PHONY: build
build: build.controller build.sidecar ## Build container images without prerequisites

.PHONY: build.controller build.sidecar
build.controller: controller/Dockerfile ## Build only the controller container image
	$(DOCKER) build --file controller/Dockerfile --platform $(PLATFORM) $(BUILD_ARGS) --tag $(CONTROLLER_TAG) .
build.sidecar: sidecar/Dockerfile ## Build only the sidecar container image
	$(DOCKER) build --file sidecar/Dockerfile --platform $(PLATFORM) $(BUILD_ARGS) --tag $(SIDECAR_TAG) .

.PHONY: generate
generate: crds controller/Dockerfile sidecar/Dockerfile ## Generate files

.PHONY: crds
crds:
	cd ./client && $(CONTROLLER_GEN) rbac:roleName=manager-role crd paths="./apis/objectstorage/..."

%/Dockerfile: hack/Dockerfile.in hack/gen-dockerfile.sh
	hack/gen-dockerfile.sh $* > "$@"

.PHONY: codegen
codegen: codegen.client codegen.proto ## Generate code

.PHONY: codegen.client codegen.proto
codegen.client:
	cd ./client && $(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./apis/objectstorage/..."
codegen.proto:
	$(MAKE) -C proto codegen

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: docs
docs: generate mdbook ## Build docs
	$(CRD_REF_DOCS) \
		--config=./docs/.crd-ref-docs.yaml \
		--source-path=./client/apis \
		--renderer=markdown \
		--output-path=./docs/src/api/
	cd docs; $(MDBOOK) build

MDBOOK_PORT ?= 3000

.PHONY: docs.serve
docs.serve: generate mdbook docs ## Serve locally built docs
	cd docs; $(MDBOOK) serve --port $(MDBOOK_PORT)

.PHONY: vendor
vendor: tidy.client tidy.proto ## Update go vendor dir
	go mod tidy
	go mod vendor
tidy.%: FORCE
	cd $* && go mod tidy

.PHONY: test-e2e
test-e2e: ## Run e2e tests against the local K8s cluster (requires both controller and driver deployed)

##@ Deployment (Advanced)

.PHONY: cluster
cluster: ## Create Kind cluster and local registry
	$(CTLPTL) apply -f ctlptl.yaml

.PHONY: cluster-reset
cluster-reset: ## Delete Kind cluster
	$(CTLPTL) delete -f ctlptl.yaml

.PHONY: deploy
deploy: ## Deploy controller (CONTROLLER_TAG) to the local K8s cluster
	./hack/dev-kustomize.sh && $(KUSTOMIZE) build $(CURDIR)/.cache | $(KUBECTL) apply -f -

.PHONY: undeploy
undeploy: ## Undeploy controller (CONTROLLER_TAG) from the local K8s cluster
	./hack/dev-kustomize.sh && $(KUSTOMIZE) build $(CURDIR)/.cache | $(KUBECTL) delete --ignore-not-found=true -f -

#
# ===== Tools =====
#

# Location to install dependencies to
TOOLBIN ?= $(CURDIR)/.cache/tools
$(TOOLBIN):
	mkdir -p $(TOOLBIN)

# Tools
GOTOOLCMD      := go tool -modfile=hack/tools/go.mod
ADDLICENSE     ?= $(GOTOOLCMD) github.com/google/addlicense
CHAINSAW       ?= $(GOTOOLCMD) github.com/kyverno/chainsaw
CRD_REF_DOCS   ?= $(GOTOOLCMD) github.com/elastic/crd-ref-docs
CTLPTL         ?= $(GOTOOLCMD) github.com/tilt-dev/ctlptl/cmd/ctlptl
GOLANGCI_LINT  ?= $(GOTOOLCMD) github.com/golangci/golangci-lint/v2/cmd/golangci-lint
KIND           ?= $(GOTOOLCMD) sigs.k8s.io/kind
KUBEAPI_LINT   ?= $(GOTOOLCMD) sigs.k8s.io/kube-api-linter/cmd/golangci-lint-kube-api-linter
KUSTOMIZE      ?= $(GOTOOLCMD) sigs.k8s.io/kustomize/kustomize/v5
LOGCHECK       ?= $(GOTOOLCMD) sigs.k8s.io/logtools/logcheck
CONTROLLER_GEN ?= $(GOTOOLCMD) sigs.k8s.io/controller-tools/cmd/controller-gen

MDBOOK         ?= $(TOOLBIN)/mdbook
SHELLCHECK     ?= $(TOOLBIN)/shellcheck

# Tool Versions
MDBOOK_VERSION           ?= v0.4.47
HADOLINT_VERSION         ?= v2.12.0
SHELLCHECK_VERSION       ?= v0.11.0

.PHONY: mdbook
mdbook: $(MDBOOK)-$(MDBOOK_VERSION)
$(MDBOOK)-$(MDBOOK_VERSION): $(TOOLBIN)
	./hack/tools/install-mdbook.sh $(MDBOOK) $(MDBOOK_VERSION)

.PHONY: shellcheck
shellcheck: $(SHELLCHECK)-$(SHELLCHECK_VERSION)
$(SHELLCHECK)-$(SHELLCHECK_VERSION): $(TOOLBIN)
	./hack/tools/install-shellcheck.sh $(SHELLCHECK) $(SHELLCHECK_VERSION)

.PHONY: FORCE # use this to force phony behavior for targets with pattern rules
FORCE:
