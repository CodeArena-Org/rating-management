build:
	@go build -o bin/main main.go

run: build
	@./bin/main

set-go-path:
	@export PATH="$PATH:$HOME/.local/bin:$(go env GOPATH)/bin"
	
grpc-gen: set-go-path
	@protoc --go_out=grpc_mngm --go_opt=paths=source_relative \
    --go-grpc_out=grpc_mngm --go-grpc_opt=paths=source_relative \
    management.proto
	
tidy:
	@go mod tidy
