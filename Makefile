BINARY_NAME=codern-user-service

build:
	GOARCH=amd64 GOOS=linux go build -o dist/${BINARY_NAME} cmd/main.go

clean:
	go clean
	rm -rf dist/${BINARY_NAME}

deps:
	go mod download

proto:
	protoc --proto_path=./protos \
	--go_out=./pkg/pb/ \
	--go-grpc_out=./pkg/pb/ \
	./protos/auth/*.proto