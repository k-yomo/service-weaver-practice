

.PHONY: setup
setup:
	go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

.PHONY: run
run:
	weaver generate .
	go run .
