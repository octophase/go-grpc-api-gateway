PROTO_FILES := $(shell find pkg -type f -name "*.proto")

.PHONY: proto
proto:
	protoc $(PROTO_FILES) --go_out=. --go-grpc_out=.

server:
	go run cmd/main.go
