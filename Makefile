# HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
# HAS_COUNTERFEITER := $(shell command -v counterfeiter;)

# #### DEPS ####
# .PHONY: deps-go-binary deps-counterfeiter deps-golangci-lint deps-modules

# deps-counterfeiter: deps-go-binary
# ifndef HAS_COUNTERFEITER
# 	go install github.com/maxbrunsfeld/counterfeiter/v6@latest
# endif

# deps-ginkgo: deps-go-binary
# ifndef HAS_GINKGO
# 	go install github.com/onsi/ginkgo/ginkgo@latest
# endif

deps-modules:
	go mod download

# #### SRC ####
# lib/libfakes/fake_firmware_store.go: lib/firmware_store.go
# 	go generate lib/firmware_store.go

# #### TEST ####
.PHONY: lint

lint:
ifndef HAS_GOLANGCI_LINT
	$(error golangci-lint is required)
endif
	golangci-lint run

# test: lib/libfakes/fake_firmware_store.go deps-modules deps-ginkgo
# 	ginkgo -r -skipPackage test .

# integration-test: deps-modules deps-ginkgo
# 	ginkgo -r test/integration

# test-all: lib/libfakes/fake_dbinterface.go deps-modules deps-ginkgo
# 	ginkgo -r .

# #### BUILD ####
.PHONY: build
SOURCES = $(shell find . -name "*.go" | grep -v "_test\." )

build/blank: $(SOURCES) deps-modules
	go build -o build/blank github.com/petewall/eink-radiator-image-source-blank/v2

build: build/blank
