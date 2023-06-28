LOCAL_BIN:=$(CURDIR)/bin

.PHONY: run
run:
	go build -o bin/ ./... && ./bin/storenum


bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.0 && \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.0 && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.3.0 && \
	go install github.com/mitchellh/gox@v1.0.1  && \
	go install golang.org/x/tools/cmd/goimports@v0.1.9 && \
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0 \

bin-local-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pav5000/smartimports/cmd/smartimports@v0.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

.PHONY: format
format: bin-local-deps
	PATH=$(LOCAL_BIN):$(PATH) smartimports -exclude 'internal/pb,pkg/api'

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint

.PHONY: lint
lint: bin-local-deps
	$(GOLANGCI_BIN) run --config=.golangci.pipeline.yaml ./...

proto-gen: proto-gen-storenum proto-gen-vendors

proto-gen-storenum:
	   find api -name "*.proto" -exec protoc -I. \
		--go_out ./pkg --go_opt paths=source_relative \
		--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./config/swagger/swaggerui \
		{} \;

proto-gen-vendors:
	find proto -name "*.proto" -exec protoc -I. \
    		--go_out ./internal/pb --go_opt paths=source_relative \
    		--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
    		--grpc-gateway_out ./internal/pb --grpc-gateway_opt paths=source_relative \
    		{} \;