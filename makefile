BINARY_NAME=go-in-tests

build:
	@GOARCH=amd64 GOOS=linux go build -ldflags '-s -w' -o ${BINARY_NAME}-linux main.go
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags '-s -w' -o ${BINARY_NAME}-darwin main.go
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags '-s -w' -o ${BINARY_NAME}-windows main.go

run:
	@go run main.go

clean:
	@go clean
	@go clean -testcache
	@if [ -e coverage.out ]; then rm coverage.out; fi;
	@if [ -e ${BINARY_NAME}-linux ]; then rm ${BINARY_NAME}-linux; fi;
	@if [ -e ${BINARY_NAME}-darwin ]; then rm ${BINARY_NAME}-darwin; fi;
	@if [ -e ${BINARY_NAME}-windows ]; then rm ${BINARY_NAME}-windows; fi;

test:
	@go test -v ./...

coverage:
	@go test ./... -coverprofile=coverage.out

lint:
	@golangci-lint run