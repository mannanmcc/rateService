package store

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/uuid"
	internalError "github.com/mannanmcc/rateService/internal/errors"
)

type Store struct {
	client         *mongo.Client
	rateCollection *mongo.Collection
}

const (
	rateCollection       = "rate_collection"
	ErrInsertingRateToDB = internalError.Error("failed to insert rate to db")
	errGettingRateFromDB = internalError.Error("failed to get rate from db")
)

var (
	nowProvider    = time.Now
	uuidProvider   = uuid.NewString
	ErrEmptyClient = internalError.Error("empty client")
	ErrEmptyDBName = errors.New("empty dbname provided")
)

func New(ctx context.Context, client *mongo.Client, dbName string) (*Store, error) {
	if client == nil {
		return nil, ErrEmptyClient
	}

	if dbName == "" {
		return nil, ErrEmptyDBName
	}

	db := client.Database(dbName)
	rateCollection := db.Collection(rateCollection)

	s := &Store{
		client:         client,
		rateCollection: rateCollection,
	}

	// if err := s.ensureIndexes(ctx); err != nil {
	// 	return nil, err
	// }

	return s, nil
}

// func (s *Store) ensureIndexes(ctx context.Context) error {
// 	_, err := s.rateCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
// 		Keys: bsonx.Doc{
// 			bsonx.Elem{Key: rateKey, Value: bsonx.String()},
// 		},
// 	})
// }
