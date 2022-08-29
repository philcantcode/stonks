package tracking

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stock struct {
	ID          primitive.ObjectID
	Ticker      string
	Company     string
	Industry    string
	OwnedStocks float64
	Group       string
	Notes       string
}
