.PHONY: protos server

protos:
	protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency

build:
	go build -o ./bin/CurrencyRate main.go

start_server:
	go run main.go server