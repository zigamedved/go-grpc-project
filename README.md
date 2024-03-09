# go-grpc-project
 
## Installation of grpc
- brew install protoc-gen-go-grpc
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
- export PATH="$PATH:$(go env GOPATH)/bin"

## Generating protobuf code
- protoc --go_out=. --go-grpc_out=. proto/greet.proto
- go mod tidy