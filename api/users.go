//go:generate protoc --gogofaster_out=plugins=grpc:. --micro_out=. --proto_path=$GOPATH/src:$GOPATH/pkg/mod:. users.proto

package users
