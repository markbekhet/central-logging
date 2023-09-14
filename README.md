# central-logging

## Commands

To compile the messages of .proto file: 
	protoc --go_out=logging-server httplog.proto
	protoc --go_out=http-server httplog.proto

To compile the services of .proto file:
	protoc --go-grpc_out=logging-server httplog.proto
