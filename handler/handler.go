package handler

import (
	"golang-starter-pack/bracelet"
)

type Handler struct {
	braceletStore bracelet.Store
}

func NewHandler(bs bracelet.Store) *Handler {
	return &Handler{
		braceletStore: bs,
	}
}
