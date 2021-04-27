.PHONY: build-v1
build-v1:
	go build -v ./...

.PHONY: test-v1
test-v1:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...