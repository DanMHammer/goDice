docs:
	go run github.com/swaggo/swag/cmd/swag@latest init

build:
	go build -o ./bin/app

PORT ?= 80

run: build
	PORT=${PORT} ./bin/app