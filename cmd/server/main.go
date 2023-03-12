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
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/mannanmcc/rateService/config"
	"github.com/mannanmcc/rateService/internal/adapter/currency"
	"github.com/mannanmcc/rateService/internal/logger"
	rateservice "github.com/mannanmcc/rateService/internal/rateservice"
	v1 "github.com/mannanmcc/rateService/internal/transport/grpc/v1"
	"github.com/mannanmcc/rateService/store"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("launce grpc application")
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Add mongodb go client to the app
	atlasURI := "mongodb://username:password@mongo/test?ssl=false&authSource=admin"
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(atlasURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	st, err := store.New(ctx, mongoClient, cfg.MongoDBName)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	//defer mongoClient.Disconnect(ctx)

	currencyProvider := currency.New(cfg.CurrencyAPIUrl, cfg.CurrencyAPIConnTimeout)
	rateService := rateservice.New(currencyProvider, st)
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
