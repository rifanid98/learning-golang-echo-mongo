package example

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct {
}

func (mc *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{
		InsertedID: nil,
	}
	return c, nil
}

var mockColl *mockCollection

func TestMain(m *testing.M) {
	mockColl = &mockCollection{}
	os.Exit(m.Run())
}

func TestInsertSuccess(t *testing.T) {
	res, err := insertData(mockColl, &User{"Adnin", "Rifandi"})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}

func TestInsertInvalidData(t *testing.T) {
	res, err := insertData(mockColl, &User{"Rifandi", "Adnin"})
	assert.NotNil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
