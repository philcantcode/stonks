package main

import (
	"fmt"

	"github.com/philcantcode/stonks/system"
	"github.com/philcantcode/stonks/yahoo"
)

func main() {
	system.InitSqlite()
	system.InitSettings()
	system.InitMongo()

	stock := yahoo.QuerySymbol("YCA.L")
	stock.INSERT_Stock()

	fmt.Println(stock)
}
