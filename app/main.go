package main

import (
	"github.com/moaabid/golang-bitly/database"
	"github.com/moaabid/golang-bitly/router"
)

func main() {
	database.ConnectDB()

	router.SetupRouter()
}
