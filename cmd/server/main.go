package main

import (
	"fmt"
	"net"
	"os"

	"log"

	"github.com/mannanmcc/rateService/config"
	"github.com/mannanmcc/rateService/internal/adapter/currency"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
	v1 "github.com/mannanmcc/rateService/internal/transport/grpc/v1"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("launce grpc application")
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	currencyProvider := currency.New(cfg.CurrencyAPIUrl)
	rateService := rateservice.New(currencyProvider)
	grpcServer := grpc.NewServer()

	handler := v1.New(rateService)
	handler.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Unavle to listen on port:", "50051")
		os.Exit(1)
	}

	//start the grpc server
	grpcServer.Serve(lis)
}
