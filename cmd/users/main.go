package main

import (
	"fmt"
	"os"

	users "github.com/koverto/users/api"
	"github.com/koverto/users/internal/pkg/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	service := micro.NewService(micro.Name("users"))
	service.Init()

	conf, err := handler.NewConfig("users", env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h, err := handler.New(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := users.RegisterUsersHandler(service.Server(), h); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
