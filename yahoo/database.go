package yahoo

import (
	"context"
	"fmt"

	"github.com/philcantcode/stonks/system"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SELECT_YFStock(filter bson.M, projection bson.M) []YFStock {
	cursor, err := system.Yahoo_Stocks_DB.Find(context.Background(), filter, options.Find().SetProjection(projection))
	system.Fatal("Couldn't SELECT_YFStock", err)
	defer cursor.Close(context.Background())

	var results []YFStock

	for cursor.Next(context.Background()) {
		var stock YFStock

		err = cursor.Decode(&stock)
		system.Fatal("Couldn't decode SELECT_YFStock", err)

		results = append(results, stock)
	}

	return results
}

func (stock YFStock) INSERT_Stock() {
	system.Log("Attempting to INSERT_YFStock", false)

	stock.ID = primitive.NewObjectID()
	insertResult, err := system.Yahoo_Stocks_DB.InsertOne(context.Background(), stock)

	system.Fatal("Couldn't INSERT_YFStock", err)
	system.Log(fmt.Sprintf("New Insert at: %s", insertResult), false)
}

func (stock YFStock) UPDATE() {
	result, err := system.Yahoo_Stocks_DB.ReplaceOne(context.Background(), bson.M{"_id": stock.ID}, stock)
	system.Fatal("Couldn't UPDATE YFStock", err)

	system.Log(fmt.Sprintf("UPDATE YFStock ID: %s, Result: %d", stock.ID, result.ModifiedCount), false)
}

func DELETE_YFStock(filter bson.M) {
	system.Log("Attempting to DELETE_YFStock", false)

	insertResult, err := system.Yahoo_Stocks_DB.DeleteOne(context.Background(), filter)

	system.Fatal("Couldn't DELETE_YFStock", err)
	system.Log(fmt.Sprintf("New Delete count: %d", insertResult.DeletedCount), false)
}
