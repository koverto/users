//go:generate protoc --gogofaster_out=plugins=grpc:. --micro_out=. --proto_path=$GOPATH/src:$GOPATH/pkg/mod:. users.proto

// Package users defines the protocol buffers API for the users service.
package users

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// AsUpdateDocument generates a MongoDB update document for a User object.
func (u *User) AsUpdateDocument() bson.M {
	currentDate := make(map[string]bool)
	set := make(map[string]interface{})

	if email := strings.TrimSpace(u.GetEmail()); email != "" {
		set["email"] = email
	}

	if name := strings.TrimSpace(u.GetName()); name != "" {
		set["name"] = name
	}

	if len(set) > 0 {
		currentDate["updatedat"] = true
	}

	return bson.M{
		"$set":         set,
		"$currentDate": currentDate,
	}
}
