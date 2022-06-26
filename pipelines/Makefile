BIN_DIR = bin
PROTO_DIR = proto

.DEFAULT_GOAL := protos
.PHONY: protos

protos:
	@echo "Generating protos..."
	@protoc -I proto --go_out=. --go-grpc_out=. ${PROTO_DIR}/*.proto
	@echo "Protos generated."

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
