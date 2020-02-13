package handler

import (
	"context"

	users "github.com/koverto/users/api"
)

type Users struct{}

func (u *Users) Create(ctx context.Context, in *users.User, out *users.User) error {
	return nil
}

func (u *Users) Get(ctx context.Context, in *users.User, out *users.User) error {
	return nil
}
