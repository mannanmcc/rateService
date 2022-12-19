package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mannanmcc/rateService/config"
	"github.com/mannanmcc/rateService/internal/adapter/currency"
	"github.com/mannanmcc/rateService/internal/logger"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
	v1 "github.com/mannanmcc/rateService/internal/transport/grpc/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("launce grpc application")
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	currencyProvider := currency.New(cfg.CurrencyAPIUrl, cfg.CurrencyAPIConnTimeout)
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
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error(context.Background(), "failed to serve grpc server", zap.Error(err))
	}

}
