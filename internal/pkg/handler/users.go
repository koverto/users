package handler

import (
	"context"
	"fmt"
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
	filter := bson.M{}

	if in.Id != nil {
		filter["_id"] = in.Id
	} else if in.Email != "" {
		filter["email"] = in.Email
	} else {
		return fmt.Errorf("no filter parameters specified")
	}

	collection := u.client.Collection("users")
	return collection.FindOne(ctx, filter).Decode(out)
}

func (u *Users) Update(ctx context.Context, in *users.User, out *users.User) error {
	return fmt.Errorf("not yet implemented") // TODO
}
