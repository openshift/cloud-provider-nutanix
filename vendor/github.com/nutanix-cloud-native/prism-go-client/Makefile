SHELL := /bin/bash
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=nutanixclient
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build

all: help

## Build:
build: ## Build your project and put the output binary in bin/
	mkdir -p bin
	$(GOCMD) build -o bin/$(BINARY_NAME) .

# go-get-tool will 'go install' any package $2 and install it to $1.
# originally copied from kubebuilder
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
BIN_PATH=$$(realpath $$(dirname $(1))) ;\
PKG_BIN_NAME=$$(echo "$(2)" | sed 's,^.*/\(.*\)@v.*$$,\1,') ;\
BIN_NAME=$$(basename $(1)) ;\
echo "Install dir $$BIN_PATH" ;\
echo "Downloading $(2)" ;\
GOBIN=$$BIN_PATH go install $(2) ;\
[[ $$PKG_BIN_NAME == $$BIN_NAME ]] || mv -f $$BIN_PATH/$$PKG_BIN_NAME $$BIN_PATH/$$BIN_NAME ;\
}
endef

TOOLS_BIN_DIR := hack/tools/bin

$(TOOLS_BIN_DIR):
	mkdir -p $(TOOLS_BIN_DIR)

CONTROLLER_GEN_BIN := controller-gen
CONTROLLER_GEN := $(TOOLS_BIN_DIR)/$(CONTROLLER_GEN_BIN)
# CRD_OPTIONS define options to add to the CONTROLLER_GEN
CRD_OPTIONS ?= "crd:crdVersions=v1"

$(CONTROLLER_GEN): $(TOOLS_BIN_DIR)
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0)

KEPLOY_VER := v0.7.12
KEPLOY_BIN := server-$(KEPLOY_VER)
KEPLOY := $(TOOLS_BIN_DIR)/$(KEPLOY_BIN)
KEPLOY_PKG := go.keploy.io/server/cmd/server

.PHONY: $(KEPLOY)
$(KEPLOY): $(TOOLS_BIN_DIR)
	$(call go-get-tool,$(KEPLOY),$(KEPLOY_PKG)@$(KEPLOY_VER))

.PHONY: run-keploy
run-keploy: $(KEPLOY)
	$(KEPLOY) run &

.PHONY: stop-keploy
stop-keploy:
	@-pkill "$(KEPLOY_BIN)"

generate: $(CONTROLLER_GEN)  ## Generate zz_generated.deepcopy.go
	$(CONTROLLER_GEN) paths="./..." object:headerFile="hack/boilerplate.go.txt"

clean: ## Remove build related file
	rm -fr ./bin vendor hack/tools/bin
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml

## Test:
test: run-keploy ## Run the tests of the project
ifeq ($(EXPORT_RESULT), true)
	go install github.com/jstemmer/go-junit-report
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
endif
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)
	@$(MAKE) stop-keploy

coverage: run-keploy ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov
ifeq ($(EXPORT_RESULT), true)
	go install github.com/AlekSi/gocov-xml
	go install github.com/axw/gocov/gocov
	gocov convert profile.cov | gocov-xml > coverage.xml
endif
	@$(MAKE) stop-keploy

## Lint:
lint: lint-go lint-yaml lint-kubebuilder ## Run all available linters

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --enable gofmt --fix --enable gofumpt $(OUTPUT_OPTIONS)

lint-yaml: ## Use yamllint on the yaml file of your projects
ifeq ($(EXPORT_RESULT), true)
	go install github.com/thomaspoignant/yamllint-checkstyle
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | yamllint-checkstyle > yamllint-checkstyle.xml)
endif
	docker run --rm -it -v $(shell pwd):/data cytopia/yamllint -c yamllint-config.yaml -f parsable $(shell git ls-files '*.yml' '*.yaml') $(OUTPUT_OPTIONS)

.PHONY: lint-kubebuilder
lint-kubebuilder: $(CONTROLLER_GEN) ## Lint Kubebuilder annotations by generating objects and checking if it is successful
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=.

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
