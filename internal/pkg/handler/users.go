package handler

import (
	"context"
	"fmt"
	"time"

	users "github.com/koverto/users/api"

	"github.com/koverto/errors"
	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const USERS_COLLECTION = "users"

type Users struct {
	*Config
	client mongo.Client
}

func New(conf *Config) (*Users, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

	var index mmongo.IndexModel
	index.Keys = bson.M{"email": 1}
	index.Options = options.Index().SetUnique(true)

	client.DefineIndexes(mongo.NewIndexSet(USERS_COLLECTION, index))
	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Users{conf, client}, nil
}

func (u *Users) Create(ctx context.Context, in *users.User, out *users.User) error {
	in.Id = uuid.New()

	now := time.Now()
	in.CreatedAt = &now

	ins, err := bson.Marshal(in)
	if err != nil {
		return err
	}

	collection := u.client.Collection(USERS_COLLECTION)
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

	collection := u.client.Collection(USERS_COLLECTION)
	err := collection.FindOne(ctx, filter).Decode(out)

	if err == mmongo.ErrNoDocuments {
		return errors.NotFound(u.ID(), "no user found: %s", filter)
	}

	return err
}

func (u *Users) Update(ctx context.Context, in *users.User, out *users.User) error {
	return errors.NotImplemented(u.ID())
}
