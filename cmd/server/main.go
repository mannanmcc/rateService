package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mannanmcc/rateService/config"
	"github.com/mannanmcc/rateService/internal/adapter/currency"
	"github.com/mannanmcc/rateService/internal/logger"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
	v1 "github.com/mannanmcc/rateService/internal/transport/grpc/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("launcing grpc application")

	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	// Add mongodb go client to the app
	atlasUri := fmt.Sprintf("mongodb://username:password@mongodb:%s/test?retryWrites=true&w=majority", cfg.MongoDbPort)
	client, err := mongo.NewClient(options.Client().ApplyURI(atlasUri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	currencyProvider := currency.New(cfg.CurrencyAPIUrl, cfg.CurrencyAPIConnTimeout)
	rateService := rateservice.New(currencyProvider)
	grpcServer := grpc.NewServer()

	handler := v1.New(rateService)
	handler.Register(grpcServer)

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal("Unable to listen on port:", "50051")
		os.Exit(1)
	}

	// start the grpc server
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error(context.Background(), "failed to serve grpc server", zap.Error(err))
	}

}
