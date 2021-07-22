package example

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func findData(colelction CollectionAPI) ([]User, error) {
	var users []User
	ctx := context.Background()

	cur, err := colelction.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("find error : %+v\n", err)
		return users, err
	}

	fmt.Printf("cursor :%+v\n", cur.Current)
	err = cur.All(ctx, &users)
	if err != nil {
		return users, err
	}
	return users, nil
}
