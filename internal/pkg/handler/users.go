package handler

import (
	"context"

	users "github.com/koverto/users/api"

	"github.com/koverto/mongo"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Users struct {
	client mongo.Client
}

func New(conf *Config) (*Users, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

	var index mmongo.IndexModel
	index.Keys = bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}}
	index.Options = options.Index().SetUnique(true)

	client.DefineIndexes(mongo.NewIndexSet("users", index))

	return &Users{
		client,
	}, nil
}

func (u *Users) Create(ctx context.Context, in *users.User, out *users.User) error {
	return nil
}

func (u *Users) Get(ctx context.Context, in *users.User, out *users.User) error {
	return nil
}
