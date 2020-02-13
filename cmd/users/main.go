package main

import (
	"fmt"
	"os"

	users "github.com/koverto/users/api"
	"github.com/koverto/users/internal/pkg/handler"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()

	if err := users.RegisterUsersHandler(service.Server(), new(handler.Users)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
