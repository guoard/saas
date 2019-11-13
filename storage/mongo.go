package storage

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "screenshots"
	collectionName = "jobs"
)

type mongoDataStore struct {
	client *mongo.Client
}

func (s *mongoDataStore) Fetch(url string) (string, error) {
	res := s.client.Database(databaseName).Collection(collectionName).FindOne(context.Background(), bson.M{
		"url": url,
	})
	if res.Err() != nil {
		return "", fmt.Errorf("failed fetching path from database: %v", res.Err())
	}
	var m bson.M
	err := res.Decode(&m)
	if err != nil {
		return "", fmt.Errorf("failed decoding object from database: %v", err)
	}
	return m["path"].(string), nil
}

func (s *mongoDataStore) FetchStatus(url string) (exists bool, isFinished bool, err error) {
	res := s.client.Database(databaseName).Collection(collectionName).FindOne(context.Background(), bson.M{
		"url": url,
	})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return false, false, nil
		}
		return false, false, fmt.Errorf("failed fetching path from database: %v", res.Err())
	}

	var m bson.M
	err = res.Decode(&m)
	if err != nil {
		return false, false, fmt.Errorf("failed decoding object from database: %v", err)
	}

	if m["status"].(string) == "finished" {
		return true, true, nil
	}

	return true, false, nil
}

func (s *mongoDataStore) Store(url string, destination string) error {
	_, err := s.client.Database(databaseName).Collection(collectionName).InsertOne(context.Background(), bson.M{
		"url":    url,
		"path":   destination,
		"status": "pending",
		"time":   time.Now().Unix(),
	})

	return err
}

func (s *mongoDataStore) SetFinished(url string) error {

	_, err := s.client.Database(databaseName).Collection(collectionName).UpdateOne(context.Background(),
		bson.M{
			"url": url,
		},
		bson.M{
			"$set": bson.M{
				"status": "finished",
			},
		})

	return err
}

func (s *mongoDataStore) SetFailed(url string) error {
	_, err := s.client.Database(databaseName).Collection(collectionName).UpdateOne(context.Background(),
		bson.M{
			"url": url,
		},
		bson.M{
			"$set": bson.M{
				"status": "failed",
			},
		})
	return err
}