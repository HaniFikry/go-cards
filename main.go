package main

import (
	"github.com/HaniFikry/go-cards/db"
	"github.com/HaniFikry/go-cards/routes"
)

func main() {
	db.Init()

	router := routes.SetupRouter()
	router.Run()
}
