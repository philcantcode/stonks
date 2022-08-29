package tracking

import (
	"context"
	"fmt"

	"github.com/philcantcode/stonks/system"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insert_stock(stock Stock) {
	system.Log("Attempting to INSERT_Capability", false)

	stock.ID = primitive.NewObjectID()
	insertResult, err := system.Tracking_Stocks_DB.InsertOne(context.Background(), stock)

	system.Fatal("Couldn't insert_stock", err)
	system.Log(fmt.Sprintf("New Insert at: %s", insertResult), false)
}

func (stock *Stock) update_stock() {
	system.Log(fmt.Sprintf("Attempting to update_stock (ID: %d)", stock.ID), true)

	result, err := system.Tracking_Stocks_DB.ReplaceOne(context.Background(), bson.M{"id": stock.ID}, stock)

	system.Fatal("Couldn't update_stock", err)
	system.Log(fmt.Sprintf("New Update made: %b", result.ModifiedCount), false)
}

func select_stock(filter bson.M, projection bson.M) []Stock {
	cursor, err := system.Tracking_Stocks_DB.Find(context.Background(), filter, options.Find().SetProjection(projection))
	system.Fatal("Couldn't select_stock", err)
	defer cursor.Close(context.Background())

	var results []Stock

	for cursor.Next(context.Background()) {
		var stock Stock

		err = cursor.Decode(&stock)
		system.Fatal("Couldn't decode select_stock: ", err)

		results = append(results, stock)
	}

	return results
}

func delete_stock(filter bson.M) {
	system.Log("Attempting to delete_stock", false)

	insertResult, err := system.Tracking_Stocks_DB.DeleteOne(context.Background(), filter)

	system.Fatal("Couldn't delete_stock", err)
	system.Log(fmt.Sprintf("New Delete count: %d", insertResult.DeletedCount), false)
}
