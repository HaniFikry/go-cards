package tests

import (
	"encoding/json"
	"github.com/HaniFikry/go-cards/models"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type DrawResponse struct {
	Cards []models.Card `json:"cards"`
}

func TestDrawCard(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	url := "/decks/" + deck.ID.String() + "/draw"

	assert.Equal(deck.Remaining, 52, "Should have 52 cards")

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 200, "Should return 200")

	var resp DrawResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	cards := resp.Cards

	assert.Equal(len(cards), 1, "Should return 1 card")
}

func TestDrawFiveCards(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	url := "/decks/" + deck.ID.String() + "/draw?count=5"

	assert.Equal(deck.Remaining, 52, "Should have 52 cards")

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 200, "Should return 200")

	var resp DrawResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	cards := resp.Cards

	assert.Equal(len(cards), 5, "Should return 1 card")

	url = "/decks/" + deck.ID.String()

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var openedDeck models.Deck
	json.Unmarshal(w.Body.Bytes(), &openedDeck)

	assert.Equal(openedDeck.Remaining, 47, "Should return the correct remaining cards")
}

func TestDrawMoreThanAvailableCards(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("POST", "/decks?cards=AS", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	assert.Equal(deck.Remaining, 1, "Should have 52 cards")

	url := "/decks/" + deck.ID.String() + "/draw?count=2"

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 400, "Should return 200")
}

func TestDrawFromEmptyDeck(t *testing.T) {
	assert := assert.New(t)
	req := httptest.NewRequest("POST", "/decks?cards=AS", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var deck models.Deck
	json.Unmarshal(w.Body.Bytes(), &deck)

	assert.Equal(deck.Remaining, 1, "Should have 52 cards")

	url := "/decks/" + deck.ID.String() + "/draw"

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	req = httptest.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(w.Code, 400, "Should return 400")
}
