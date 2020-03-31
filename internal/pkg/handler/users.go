// Package handler defines the gRPC endpoint handlers for the Users service.
package handler

import (
	"context"
	"fmt"
	"time"

	users "github.com/koverto/users/api"

	"github.com/koverto/errors"
	"github.com/koverto/micro/v2"
	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const usersCollection = "users"

// Users defines the Users service.
type Users struct {
	*Config
	*micro.Service
	client mongo.Client
}

// Config contains the configuration for an instance of the Users service
// handlers.
type Config struct {
	MongoURL string `json:"mongourl"`
}

// New creates a new instance of the Users service handlers.
func New(conf *Config, service *micro.Service) (*Users, error) {
	client, err := mongo.NewClient(conf.MongoURL, "users")
	if err != nil {
		return nil, err
	}

	var index mmongo.IndexModel
	index.Keys = bson.M{"email": 1}
	index.Options = options.Index().SetUnique(true)

	client.DefineIndexes(mongo.NewIndexSet(usersCollection, index))

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Users{conf, service, client}, nil
}

// Create inserts a new User object into the database.
func (u *Users) Create(ctx context.Context, in *users.User, out *users.User) error {
	in.Id = uuid.New()

	now := time.Now()
	in.CreatedAt = &now

	ins, err := bson.Marshal(in)
	if err != nil {
		return err
	}

	collection := u.client.Collection(usersCollection)
	_, err = collection.InsertOne(ctx, ins)

	if err == nil {
		*out = *in
	}

	return err
}

// Read gets a User object from the database by ID or email address.
func (u *Users) Read(ctx context.Context, in *users.User, out *users.User) error {
	filter := bson.M{}

	switch {
	case in.Id != nil:
		filter["_id"] = in.GetId()
	case in.Email != "":
		filter["email"] = in.GetEmail()
	default:
		return fmt.Errorf("no filter parameters specified")
	}

	collection := u.client.Collection(usersCollection)
	err := collection.FindOne(ctx, filter).Decode(out)

	if err == mmongo.ErrNoDocuments {
		return errors.NotFound(u.Name, "no user found: %s", filter)
	}

	return err
}

// Update updates a User object stored in the database.
func (u *Users) Update(ctx context.Context, in *users.User, out *users.User) error {
	update := in.AsUpdateDocument()
	if len(update) == 0 {
		return nil
	}

	filter := bson.M{"_id": in.GetId()}
	collection := u.client.Collection(usersCollection)
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	return collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(out)
}
