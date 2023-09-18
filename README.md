# central-logging


## Environment

The logging server communicates with a PostgresSQL database.

Information about the connection to the DB are retreived from variable
environments

Four variable environments are retreived from the os and if any one not found it is rendred to empty string.
1- dbHost // This will allow scalability.
2- dbUser
3- dbPassword


## Commands

To compile the messages of .proto file: 
	protoc --go_out=logging-server httplog.proto
	protoc --go_out=http-server httplog.proto

To compile the services of .proto file:
	protoc --go-grpc_out=logging-server httplog.proto
	protoc --go-grpc_out=http-server httplog.proto
