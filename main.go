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

	// ObjectID = 60f70baf04a2c03f2c464799
	// made up of 12 bytes
	// contains information about timestamp, machine id, process id, counter

	// BSON (Binary Encoded Json), includes additional types e.g. int, long, date, floating point.
	// bson.D = ordered document bson.D{{"hello", "world"}}
	// bson.M = unordered document/map bson.M{"hello": "world"}
	// bson.A = array bson.A{"element 1", "element 2"}
	// bson.E = usually used as an element inside bson.D

	db := client.Database("tronics")
	collection := db.Collection("products")

	// res, err := collection.InsertOne(context.Background(), trimmer)
	// res, err := collection.InsertOne(context.Background(), bson.D{
	//   {"name", "eric"},
	//   {"surname", "cartman"},
	//   {"hobbies", bson.A{"videogame", "alexa", "kfc"}},
	// })
	// res, err := collection.InsertOne(context.Background(), bson.M{
	//   {"name": "eric"},
	//   {"surname": "cartman"},
	//   {"hobbies": bson.A{"videogame", "alexa", "kfc"}},
	// })

	res, err := collection.InsertOne(context.Background(), iphone10)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}
