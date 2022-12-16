package main

import (
	"fmt"
	"net"
	"os"

	"log"

	protos "github.com/mannanmcc/proto/rates/rate"
	"github.com/mannanmcc/rateService/config"
	"github.com/mannanmcc/rateService/internal/adapter/currency"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
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

	//register the rateservice server
	protos.RegisterRateServiceServer(grpcServer, rateService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Unavle to listen on port:", "50051")
		os.Exit(1)
	}

	//start the grpc server
	grpcServer.Serve(lis)
}
