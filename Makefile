# Flashwire Makefile

# Version and tool locations
PROTOC_VERSION=21.12
TOOLS_DIR=.tools
PROTOC_DIR=$(TOOLS_DIR)/protoc

# Dynamic type (default int32 if not provided)
TYPE?=int32

# Directories based on TYPE
PROTO_DIR=internal/$(TYPE)/proto
BENCHMARK_DIR=internal/$(TYPE)/codec/benchmark

# OS detection
OS := $(shell uname -s)
ARCH := $(shell uname -m)

# Choose protoc zip file based on OS and Arch
PROTOC_ZIP :=
ifeq ($(OS),Linux)
	PROTOC_ZIP=protoc-$(PROTOC_VERSION)-linux-x86_64.zip
endif
ifeq ($(OS),Darwin)
	ifeq ($(ARCH),x86_64)
		PROTOC_ZIP=protoc-$(PROTOC_VERSION)-osx-x86_64.zip
	endif
	ifeq ($(ARCH),arm64)
		PROTOC_ZIP=protoc-$(PROTOC_VERSION)-osx-aarch_64.zip
	endif
endif

# Target: Install protoc-gen-go plugin globally (one-time)
install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Target: Download protoc binary locally into .tools/protoc
download-protoc:
	@echo "Detected OS: $(OS), Architecture: $(ARCH)"
	@if [ -z "$(PROTOC_ZIP)" ]; then echo "Unsupported OS/Arch combination"; exit 1; fi
	mkdir -p $(PROTOC_DIR)
	curl -L https://github.com/protocolbuffers/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_ZIP) -o $(PROTOC_DIR)/protoc.zip
	unzip -o $(PROTOC_DIR)/protoc.zip -d $(PROTOC_DIR)

# Target: Generate Go code from .proto files
proto:
	$(PROTOC_DIR)/bin/protoc --proto_path=. --go_out=. $(PROTO_DIR)/testint32.proto

gen-proto50:
	$(PROTOC_DIR)/bin/protoc --proto_path=. --go_out=. $(PROTO_DIR)/testint32_50.proto

# Benchmarks
bench-flashwire:
	go test -bench=BenchmarkFlashwire -benchmem ./$(BENCHMARK_DIR)

bench-protobuf:
	go test -bench=BenchmarkProtobuf -benchmem ./$(BENCHMARK_DIR)

bench-all: bench-flashwire bench-protobuf

# Unit tests across all packages
test:
	go test -v ./...

bench:
	go test -bench=Benchmark -benchmem ./...

fuzz:
	go test -fuzz=Fuzz -fuzztime=30s ./...

gen-datainput-int32:
	cd $(CURDIR) && go run internal/int32/generator-demo/datainput/main.go

# Not optimum, mainly due to multiple method invocations
gen-datainput50-int32:
	cd $(CURDIR) && go run internal/int32/generator-demo/datainput50/main.go
