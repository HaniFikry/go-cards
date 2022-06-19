package tests

import (
	"encoding/json"
	"github.com/HaniFikry/go-cards/models"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestOpenDeckNotExisting(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/decks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 404, "Should return 404")
}

func TestOpenDeckExisting(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	url := "/decks/" + deck.ID.String()

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var openedDeck models.Deck
	json.Unmarshal(w.Body.Bytes(), &openedDeck)

	var cards []models.Card
	json.Unmarshal([]byte(openedDeck.Cards), &cards)

	assert.Equal(openedDeck.ID, deck.ID, "Should return the same ID")
	assert.Equal(openedDeck.Remaining, deck.Remaining, "Should return the same remaining")
	assert.Equal(openedDeck.Shuffled, deck.Shuffled, "Should return the same shuffled")
	assert.NotNil(openedDeck.Cards, "Should have Cards")
	assert.Equal(len(cards), 52, "Should return 52 cards")
}
