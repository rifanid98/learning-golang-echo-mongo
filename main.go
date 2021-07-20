package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product describes an electronic product e.g. phone
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

var iphone10 = &Product{
	ID:          primitive.NewObjectID(),
	Name:        "iPhone 10",
	Price:       900,
	Currency:    "USD",
	Quantity:    "40",
	Vendor:      "Apple",
	Accessories: []string{"charger", "earphone", "slotopener"},
	SkuID:       "1234",
}

var trimmer = &Product{
	ID:          primitive.NewObjectID(),
	Name:        "Easy Trimmer",
	Price:       120,
	Currency:    "USD",
	Quantity:    "300",
	Vendor:      "Philips",
	Discount:    7,
	Accessories: []string{"charger", "comb", "bladeset", "cleaning oil"},
	SkuID:       "2345",
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	db := client.Database("tronics")
	collection := db.Collection("products")

	res, err := collection.InsertOne(context.Background(), iphone10)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
