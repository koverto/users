package users

import (
	"github.com/micro/go-micro/v2/client"
)

const Name = "com.koverto.svc.users"

type Client struct {
	UsersService
}

func NewClient(client client.Client) *Client {
	return &Client{NewUsersService(Name, client)}
}

func (c *Client) Name() string {
	return Name
}
