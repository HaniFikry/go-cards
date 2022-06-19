package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"math/rand"
	"strings"
)

type Deck struct {
	ID        uuid.UUID      `gorm:"column:deck_id;type:char(36);primary_key" json:"deck_id"`
	Shuffled  bool           `gorm:"default:false" json:"shuffled"`
	Remaining int            `json:"remaining"`
	Cards     datatypes.JSON `json:"cards"`
}

type DeckWithoutCards struct {
	ID        uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
}

func (d *Deck) Shuffle() {
	var cards []Card
	json.Unmarshal(d.Cards, &cards)
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	json_cards, _ := json.Marshal(cards)
	d.Cards = datatypes.JSON([]byte(json_cards))
}

func CreatePartialDeck(cardsToCreate string) ([]Card, bool) {
	cardsArray := strings.Split(cardsToCreate, ",")
	colorsDict := map[string]string{"S": "Spades", "D": "Diamonds", "C": "Clubs", "H": "Hearts"}
	valuesDict := map[string]string{"A": "Ace", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "10": "10", "J": "Jack", "Q": "Queen", "K": "King"}
	allCards := []Card{}
	invalid := false
	for _, card := range cardsArray {
		valueInitial := string(card[0 : len(card)-1])
		colorInitial := string(card[len(card)-1:])
		if colorsDict[colorInitial] == "" || valuesDict[valueInitial] == "" {
			invalid = true
			break
		}
		allCards = append(allCards, Card{Value: valuesDict[valueInitial], Suit: colorsDict[colorInitial], Code: card})
	}
	return allCards, invalid
}

func CreateFullDeck() []Card {
	colors := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	colorInitials := []string{"S", "D", "C", "H"}
	valueInitials := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	allCards := []Card{}
	for i, color := range colors {
		for j, value := range values {
			valueInitial := valueInitials[j]
			colorInitial := colorInitials[i]
			code := valueInitial + colorInitial
			allCards = append(allCards, Card{Value: value, Suit: color, Code: code})
		}
	}
	return allCards
}
