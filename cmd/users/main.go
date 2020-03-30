package main

import (
	users "github.com/koverto/users/api"
	"github.com/koverto/users/internal/pkg/handler"

	"github.com/koverto/micro"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := &handler.Config{
		MongoUrl: "mongodb://localhost:27017",
	}

	service, err := micro.NewService("com.koverto.svc.users", conf, env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("could not initialize service")
	}

	h, err := handler.New(conf, service)
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("could not build handler")
	}

	if err := users.RegisterUsersHandler(service.Server(), h); err != nil {
		log.Fatal().AnErr("error", err).Msg("could not register handler with service")
	}

	if err := service.Run(); err != nil {
		log.Fatal().AnErr("error", err).Msg("error running service")
	}
}
