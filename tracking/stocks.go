package tracking

import (
	"github.com/philcantcode/stonks/system"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTracked() []Stock {
	return select_stock(bson.M{}, bson.M{})
}

func TrackNewTicker(ticker string) {
	stock := Stock{Ticker: ticker}

	for _, existingStock := range GetTracked() {
		if ticker == existingStock.Ticker {
			return
		}
	}

	system.Log("Adding new ticker: "+ticker, true)

	insert_stock(stock)
}
