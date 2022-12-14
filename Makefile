.PHONY: build
build:
			 go build -v ./cmd/apiserver

.PHONY: test
test:
			 go test -v -race -timeout 30s ./...| sed "/PASS/s//$$(printf "\033[32mPASS\033[0m")/" | sed "/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/" | sed "/RUN/s//$$(printf "\033[34mRUN\033[0m")/"

DEFAULT_GOAL := build