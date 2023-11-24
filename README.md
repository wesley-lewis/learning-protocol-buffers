
### Compiling .proto files
protoc --go_out=. --go_opt=paths=source_relative ./path/to/.proto/file

## Compiling .proto files with services
protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative filename.proto
