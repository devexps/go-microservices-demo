# Not Support Windows

.PHONY: help wire gen ent build api openapi init all

CURRENT_DIR := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

SRCS_MK := $(foreach dir, app, $(wildcard $(dir)/*/*/Makefile))

# initialize develop environment
init:
	@which micro &> /dev/null || go install github.com/devexps/go-micro/cmd/micro/v2@latest
	@which protoc-gen-go-http &> /dev/null || go install github.com/devexps/go-micro/cmd/protoc-gen-go-http/v2@latest
	@which protoc-gen-go-errors &> /dev/null || go install github.com/devexps/go-micro/cmd/protoc-gen-go-errors/v2@latest

	@which protoc-gen-go &> /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	@which protoc-gen-validate &> /dev/null || go install github.com/envoyproxy/protoc-gen-validate@v1.0.0
	@which protoc-gen-go-grpc &> /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@which protoc-gen-openapi &> /dev/null || go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@which wire &> /dev/null || go install github.com/google/wire/cmd/wire@latest
	@which golangci-lint &> /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest


# generate wire code
wire:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make wire;\
    )

# generate code
gen:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make gen;\
    )

# generate ent code
ent:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make ent;\
    )

# generate protobuf api go code
api:
	buf generate

# generate OpenAPI v3 docs.
openapi:
	buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml

# build all service applications
build:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make build;\
    )

# generate & build all service applications
all:
	$(foreach dir, $(dir $(realpath $(SRCS_MK))),\
      cd $(dir);\
      make app;\
    )

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
