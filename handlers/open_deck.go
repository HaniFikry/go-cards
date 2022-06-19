package handlers

import (
	"errors"
	"github.com/HaniFikry/go-cards/db"
	"github.com/HaniFikry/go-cards/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func OpenDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deck models.Deck
		DB := db.GetDB()
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Deck id is required"})
			return
		}
		err := DB.First(&deck, "deck_id = ?", id).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Deck not found"})
			return
		}
		c.JSON(http.StatusOK, deck)
	}
}
