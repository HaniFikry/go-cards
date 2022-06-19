package tests

import (
	"encoding/json"
	"github.com/HaniFikry/go-cards/models"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCreateDeckNotShuffled(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 201, "Should return 201")

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	assert.Equal(deck.Remaining, 52, "Should return 52 cards")
	assert.Equal(deck.Shuffled, false, "Should not be shuffled")
	assert.NotNil(deck.ID, "Should have an ID")
}

func TestCreateDeckShuffled(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks?shuffled=true", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 201, "Should return 201")

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	assert.Equal(deck.Remaining, 52, "Should return 52 cards")
	assert.Equal(deck.Shuffled, true, "Should be shuffled")
	assert.NotNil(deck.ID, "Should have an ID")
}

func TestCreatePartialDeck(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks?cards=AS,5D,9H", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 201, "Should return 201")

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	assert.Equal(deck.Remaining, 3, "Should return 3 cards")
	assert.Equal(deck.Shuffled, false, "Should be shuffled")
	assert.NotNil(deck.ID, "Should have an ID")
}

func TestCreatePartialDeckWithWrongCards(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks?cards=HI", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 400, "Should return 400")
}
