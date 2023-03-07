

.PHONY: setup
setup:
	go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

.PHONY: build
build:
	make generate
	go build -o bin/server .

.PHONY: run
run:
	make generate
	SERVICEWEAVER_CONFIG=weaver.toml go run .

.PHONY: run-multi-process
run-multi-process:
	make build
	weaver multi deploy weaver.toml

.PHONY: generate
generate:
	weaver generate .
