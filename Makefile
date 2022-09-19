all:  run

run-server:
	bin/server

run-client:
	bin/client

build: build-server build-client

build-server:
	go build -o bin/server cmd/server/main.go

build-client:
	go build -o bin/client cmd/client/main.go

lint:
	golangci-lint run ./...

test:
	go test ./...

