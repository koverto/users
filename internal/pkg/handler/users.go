package handler

import (
	"context"
	"time"

	users "github.com/koverto/users/api"

	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
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
	in.Id = uuid.New()
	*in.CreatedAt = time.Now()

	ins, err := bson.Marshal(in)
	if err != nil {
		return err
	}

	collection := u.client.Collection("users")
	_, err = collection.InsertOne(ctx, ins)

	if err == nil {
		*out = *in
	}

	return err
}

func (u *Users) Read(ctx context.Context, in *users.User, out *users.User) error {
	return nil
}

func (u *Users) Update(ctx context.Context, in *users.User, out *users.User) error {
	return nil
}
