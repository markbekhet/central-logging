# central-logging


#Requirements
1- Install proto gRPC. On Ubuntu: 
sudo apt  install protobuf-compiler

2- Install go plugins for grpc:
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/golang/protobuf/protoc-gen-go@latest

## Environment
In the each of the following directories create a vendor directory and execute:
```go mod tidy```

The logging server communicates with a PostgresSQL database.

Information about the connection to the DB are retreived from variable
environments

Four variable environments are retreived from the os and if any one not found it is rendred to empty string.
1- dbHost // This will allow scalability.
2- dbUser
3- dbPassword


## Commands

To compile the messages of .proto file: 
	protoc --go_out=. httplog.proto

To compile the services of .proto file:
	protoc --go-grpc_out=. httplog.proto

Note: You may need to provide the --plugin flag to specify the path to the go-grpc plugins if the plugins are not installed in the GOROOT
