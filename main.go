package main

import (
	"fmt"
	"os"

	"github.com/philcantcode/stonks/system"
	"github.com/philcantcode/stonks/yahoo"
)

func main() {
	system.InitSqlite()
	system.InitSettings()
	system.InitMongo()

	fmt.Println(yahoo.QuerySymbol("AAPL"))
	os.Exit(0)

	initServer()
}
