

.PHONY: setup
setup:
	go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

.PHONY: run
run:
	weaver generate .
	go run .

.PHONY: run-multi-process
run-multi-process:
	make build
	weaver multi deploy weaver.toml

.PHONY: build
build:
	go build -o bin/server .

