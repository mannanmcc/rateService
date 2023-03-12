package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Rate struct {
	ID        string    `bson:"_id"`
	RateValue float32   `bson:"rate_value"`
	RateKey   string    `bson:"rate_key"`
	AddedOn   time.Time `bson:"added_on"`
}

func (s Store) AddRate(ctx context.Context, rate float32, key string) (*Rate, error) {
	r := &Rate{
		RateValue: rate,
		RateKey:   key,
	}
	r.ID = uuidProvider()
	r.AddedOn = nowProvider().UTC()
	_, err := s.rateCollection.InsertOne(ctx, r)
	if err != nil {
		return nil, ErrInsertingRateToDB.Wrap(err)
	}

	return r, nil
}

func (s Store) GetRate(ctx context.Context, key string) (Rate, error) {
	var rateFromDB Rate
	err := s.rateCollection.FindOne(ctx, bson.M{"RateKey": key}).Decode(&rateFromDB)
	if err != nil {
		return Rate{}, errGettingRateFromDB.Wrap(err)
	}

	return rateFromDB, nil
}
