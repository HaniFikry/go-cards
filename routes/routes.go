package routes

import (
	"github.com/HaniFikry/go-cards/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/decks", handlers.CreateDeck())
	router.GET("/decks/:id", handlers.OpenDeck())
	router.GET("/decks/:id/draw", handlers.DrawCard())

	return router
}
