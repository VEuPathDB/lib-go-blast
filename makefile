.PHONY: build
build: build-v1 build-v2

.PHONY: test
test: test-v1 test-v2

.PHONY: build-v1
build-v1:
	@go build -v ./v1/...

.PHONY: test-v1
test-v1:
	@echo Running v1 tests...
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v ./v1/...

.PHONY: build-v2
build-v2:
	@cd v2 && go build -v ./...

.PHONY: test-v2
test-v2:
	@echo Running v2 tests...
	@cd v2 && go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
