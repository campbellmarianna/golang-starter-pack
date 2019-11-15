package handler

import (
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"golang-starter-pack/model"
)


type braceletCreateRequest struct {
	Bracelet struct {
		Text       string   `json:"text" validate:"required"`
		ThreadColor string   `json:"threadColor" validate:"required"`
		Font        string   `json:"font" validate:"required"`
	} `json:"bracelet"`
}

func (r *braceletCreateRequest) bind(c echo.Context, a *model.Bracelet) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Text = r.Bracelet.Text
	a.Slug = slug.Make(r.Bracelet.Text)
	a.ThreadColor = r.Bracelet.ThreadColor
	a.Font = r.Bracelet.Font
	return nil
}

type createBeadRequest struct {
	Bead struct {
		Color string `json:"color" validate:"required"`
	} `json:"bead"`
}

func (r *createBeadRequest) bind(c echo.Context, bm *model.Bead) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	bm.Color = r.Bead.Color
	return nil
}
