package main

import (
	"github.com/philcantcode/stonks/system"
)

func main() {
	system.InitSqlite()
	system.InitSettings()
	system.InitMongo()

	initServer()
}
