package tracking

import (
	"strconv"

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

func UpdateTracking(ticker, company, industry, ownedStock, group, notes string) {
	tracking := GetTracked()

	for _, stock := range tracking {
		if stock.Ticker == ticker {
			stock.Company = company
			stock.Industry = industry
			stock.Group = group
			stock.Notes = notes

			if convOwnedStock, err := strconv.ParseFloat(ownedStock, 64); err == nil {
				stock.OwnedStocks = convOwnedStock
			}
		}

		stock.update_stock()
	}
}
