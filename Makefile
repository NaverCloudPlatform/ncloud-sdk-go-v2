GOFMT_FILES?=$$(find . -not -path "./vendor/*" -name "*.go")

ifeq ($(OS)),Windows_NT)
	SDK_ONLY_PKGS=$(shell go list ./... | findstr /v "\/vendor")
else
	SDK_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/")
endif

all: build

test:
	@echo "go(ginkgo) test SDK"
	@ginkgo -r

fmt:
	@gofmt -w -s $(GOFMT_FILES)

build:
	@echo "go build SDK"
	@go build ${SDK_ONLY_PKGS}
