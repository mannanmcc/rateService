package main

import (
	"fmt"
	"net"
	"os"

	"log"

	protos "github.com/mannanmcc/proto/rates/rate"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("launce grpc application")
	rateService := rateservice.New()
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
