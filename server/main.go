package main

import (
	"github.com/philcantcode/stonks/server/system"
)

func main() {
	system.InitSqlite()
	system.InitSettings()
	system.InitMongo()

	initServer()
}
