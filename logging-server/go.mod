module github.com/markbekhet/central-logging/logging-server

go 1.21.1

replace github.com/markbekhet/central-logging/log => ../log

require (
	google.golang.org/grpc v1.61.0 // direct
	google.golang.org/protobuf v1.32.0 // indirect; direct
)

require (
	github.com/lib/pq v1.10.9
	github.com/markbekhet/central-logging/log v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
)
