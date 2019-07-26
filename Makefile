# Makefile

include .env

.EXPORT_ALL_VARIABLES:	

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/ttn-service-integration

dev:
	go run main.go

prod: build
	upx ./bin/ttn-service-integration
