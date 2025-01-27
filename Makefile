.PHONY: rundev
rundev:
	@go build -o bin/gedebox cmd/
	@go run bin/gedebox

.PHONY: build
build:
	@bash build.sh

.PHONY: test
test:
	@go test ./...

.PHONY: testv
testv:
	@go test ./... -v