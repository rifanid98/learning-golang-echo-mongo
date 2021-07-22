package example

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// User a dummy user
type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

// actual collection *mongon.Collection
// implements CollectionAPI
// InsertOne will work on actual collection

// fake collection mockCollection
// implements CollectionAPI
// InsertOne fake implementation, when invoked it will works on fake collection

func insertData(collection CollectionAPI, user *User) (*mongo.InsertOneResult, error) {
	if user.FirstName != "Adnin" {
		return nil, errors.New("invalid first name")
	}

	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return res, err
	}
	return res, nil
}
