package handler

import (
	"context"

	users "github.com/koverto/users/api"

	"github.com/koverto/mongo"
)

type Users struct {
	client mongo.Client
}

func New(conf *Config) (*Users, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

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
