.DEFAULT_GOAL := run

.PHONY: build
build: gen
	go build ./cmd/users

.PHONY: docker
docker: build
	docker build . -t koverto/users:latest

.PHONY: gen
gen:
	go generate ./api

.PHONY: run
run: gen
	go run ./cmd/users
