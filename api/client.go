package users

import (
	"github.com/micro/go-micro/v2/client"
)

// Name is the identifying name of the Users service.
const Name = "com.koverto.svc.users"

// Client defines a client for the Users service.
type Client struct {
	UsersService
}

// NewClient creates a new client for the Users service.
func NewClient(client client.Client) *Client {
	return &Client{NewUsersService(Name, client)}
}

// Name returns the name of the Users service.
func (c *Client) Name() string {
	return Name
}
