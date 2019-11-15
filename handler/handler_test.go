package handler

import (
	"log"
	"os"
	"testing"

	"encoding/json"

	"golang-starter-pack/bracelet"
	"golang-starter-pack/db"
	"golang-starter-pack/model"
	"golang-starter-pack/router"
	"golang-starter-pack/store"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
)

var (
	d  *gorm.DB
	us user.Store
	bs bracelet.Store
	h  *Handler
	e  *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	d = db.TestDB()
	db.AutoMigrate(d)
	bs = store.NewBraceletStore(d)
	h = NewHandler(bs)
	e = router.New()
	loadFixtures()
}

func tearDown() {
	_ = d.Close()
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m[key].(map[string]interface{})
}

func loadFixtures() error {
	b := model.Bracelet{
		Slug:        "Black-Girls-Code-slug",
		Text:       "Black Girls Code text",
		ThreadColor: "blue threadColor",
		Font:        "bold playful font",
	}

	bs.CreateBracelet(&b)
	bs.AddBead(&b, &model.Bead{
		Color:      "pink",
		BraceletID: 1,
	})

	b2 := model.Bracelet{
		Slug:        "Naturally-MEE-slug",
		Text:       "Naturally MEE text",
		ThreadColor: "black threadColor",
		Font:        "curly",
	}
	bs.CreateBracelet(&b2)
	bs.AddBead(&b2, &model.Bead{
		Color:      "green",
		BraceletID: 2,
	})

	return nil
}
