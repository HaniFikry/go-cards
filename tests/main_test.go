package tests

import (
	"github.com/HaniFikry/go-cards/db"
	"github.com/HaniFikry/go-cards/routes"
	"os"
	"testing"
)

var router = routes.SetupRouter()

func TestMain(m *testing.M) {
	// initialize db
	db.Init()
	// Run the tests
	code := m.Run()
	//clean up test db
	os.Remove("cards.db")
	os.Exit(code)
}
