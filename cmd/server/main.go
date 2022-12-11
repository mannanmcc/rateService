package main

import (
	"fmt"
	"net"
	"os"

	"log"

	protos "github.com/mannanmcc/proto/rates/rate"
	"github.com/mannanmcc/rateService/internal/adapter/currency"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("launce grpc application")
	currencyProvider := currency.New("https://api.exchangerate.host/latest")
	rateService := rateservice.New(currencyProvider)
	grpcServer := grpc.NewServer()

	//register the rateservice server
	protos.RegisterRateServiceServer(grpcServer, rateService)

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("Unavle to listen on port:", "9001")
		os.Exit(1)
	}

	//start the grpc server
	grpcServer.Serve(lis)
}
