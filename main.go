package main

import (
	"golang-starter-pack/db"
	"golang-starter-pack/handler"
	"golang-starter-pack/router"
	"golang-starter-pack/store"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	bs := store.NewBraceletStore(d)
	h := handler.NewHandler(bs)
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
