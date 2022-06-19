package handlers

import (
	"encoding/json"
	"errors"
	"github.com/HaniFikry/go-cards/db"
	"github.com/HaniFikry/go-cards/models"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DrawCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deck models.Deck
		DB := db.GetDB()
		id := c.Param("id")
		count := c.DefaultQuery("count", "1")
		err := DB.First(&deck, "deck_id = ?", id).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Deck not found"})
			return
		}
		intCount, _ := strconv.Atoi(count)
		if deck.Remaining == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Deck is empty"})
			return
		}
		if deck.Remaining < intCount {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Remaining cards are less than the requseted count"})
			return
		}
		var cards []models.Card
		json.Unmarshal(deck.Cards, &cards)
		card := cards[0:intCount]
		cards = cards[intCount:]
		json_cards, _ := json.Marshal(cards)
		deck.Cards = datatypes.JSON([]byte(json_cards))
		deck.Remaining = deck.Remaining - intCount
		DB.Save(&deck)
		c.JSON(http.StatusOK, gin.H{
			"cards": card,
		})
	}
}
