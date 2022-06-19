package handlers

import (
	"encoding/json"
	"github.com/HaniFikry/go-cards/db"
	"github.com/HaniFikry/go-cards/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"net/http"
	"strconv"
)

func CreateDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		DB := db.GetDB()
		shuffled := c.DefaultQuery("shuffled", "false")
		cardsToCreate := c.DefaultQuery("cards", "")

		shuffledBool, _ := strconv.ParseBool(shuffled)

		cards := []models.Card{}
		invalid := false

		if cardsToCreate != "" {
			cards, invalid = models.CreatePartialDeck(cardsToCreate)
		} else {
			cards = models.CreateFullDeck()
		}

		if invalid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card"})
			return
		}

		json_cards, _ := json.Marshal(cards)
		newDeck := models.Deck{ID: uuid.New(), Remaining: len(cards), Shuffled: shuffledBool, Cards: datatypes.JSON([]byte(json_cards))}
		if shuffledBool {
			newDeck.Shuffle()
		}
		DB.Create(&newDeck)
		c.JSON(http.StatusCreated, models.DeckWithoutCards{ID: newDeck.ID, Shuffled: newDeck.Shuffled, Remaining: newDeck.Remaining})
	}
}
