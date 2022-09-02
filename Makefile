.PHONY: build example lint test clean help

build:
	@go build ./...
	@mkdir -p dist

example:
	@go build -o dist/ ./examples/form
	@go build -o dist/ ./examples/layout
	@go build -o dist/ ./examples/menu
	@go build -o dist/ ./examples/tabview
	@go build -o dist/ ./examples/webview
	@go build -o dist/ ./examples/widgets
	@go build -o dist/ ./examples/webshot

lint:
	@go fmt ./
	@go vet ./

test:
	@go test ./...

clean:
	@go clean
	@rm -rf dist/

help:
	@echo "make - format and build binary output"
	@echo "make build - build binary output"
	@echo "make clean - remove all  build outputs"
	@echo "make lint - go lint using 'fmt' and 'vet'"%