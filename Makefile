.PHONY: generate, run

generate:
	go run github.com/99designs/gqlgen generate

run:
	go run server.go