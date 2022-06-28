BIN_DIR = bin
PROTO_PIPELINES_DIR = pipelines
PROTO_UI_DIR = ui/proto
PROTO_UI_V2_DIR = ui-v2/proto

.PHONY: all

all: protos-pipelines protos-ui protos-ui-v2

protos-pipelines:
	@echo "Generating protos..."
	if not exist "$(PROTO_PIPELINES_DIR)" mkdir "$(PROTO_PIPELINES_DIR)"
	@protoc -I _proto _proto/*.proto --proto_path=./_proto --go_out=./${PROTO_PIPELINES_DIR} --go-grpc_out=./${PROTO_PIPELINES_DIR}
	@echo "Protos generated."
protos-ui:
	@echo "Generating protos..."
	if not exist "$(PROTO_UI_DIR)" mkdir "$(PROTO_UI_DIR)"
	@protoc -I _proto _proto/*.proto --proto_path=./_proto --js_out=import_style=commonjs:./${PROTO_UI_DIR} --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./${PROTO_UI_DIR}
	@echo "Protos generated."
protos-ui-v2:
	@echo "Generating protos..."
	if not exist "$(PROTO_UI_V2_DIR)" mkdir "$(PROTO_UI_V2_DIR)"
	@protoc -I _proto _proto/*.proto --proto_path=./_proto --js_out=import_style=commonjs:./${PROTO_UI_V2_DIR} --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./${PROTO_UI_V2_DIR}
	@echo "Protos generated."

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
