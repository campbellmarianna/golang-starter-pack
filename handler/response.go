package handler

import (
	"golang-starter-pack/model"
	"github.com/labstack/echo/v4"
)

type braceletResponse struct {
	Slug           string    `json:"slug"`
	Text          string    `json:"text"`
	ThreadColor    string    `json:"threadColor"`
	Font           string    `json:"font"`
}

type singleBraceletResponse struct {
	Bracelet *braceletResponse `json:"bracelet"`
}

type braceletListResponse struct {
	Bracelets      []*braceletResponse `json:"bracelets"`
	BraceletsCount int                `json:"braceletsCount"`
}

func newBraceletResponse(c echo.Context, b *model.Bracelet) *singleBraceletResponse {
	br := new(braceletResponse)
	br.Slug = b.Slug
	br.Text = b.Text
	br.ThreadColor = b.ThreadColor
	br.Font = b.Font
	return &singleBraceletResponse{br}
}

func newBraceletListResponse(bracelets []model.Bracelet, count int) *braceletListResponse {
	r := new(braceletListResponse)
	r.Bracelets = make([]*braceletResponse, 0)
	for _, b := range bracelets {
		br := new(braceletResponse)
	  br.Slug = b.Slug
	  br.Text = b.Text
	  br.ThreadColor = b.ThreadColor
	  br.Font = b.Font
		r.Bracelets = append(r.Bracelets, br)
	}
	r.BraceletsCount = count
	return r
}

type beadResponse struct {
	ID        uint      `json:"id"`
	Color      string    `json:"color"`
}

type singleBeadResponse struct {
	Bead *beadResponse `json:"bead"`
}

type beadListResponse struct {
	Beads []beadResponse `json:"beads"`
}
// These two function may through errors I did ctrl + shift + f as well as find and replace and I'm unsure if updated correctly
func newBeadResponse(c echo.Context, bm *model.Bead) *singleBeadResponse {
	bead := new(beadResponse)
	bead.ID = bm.ID
	bead.Color = bm.Color
	return &singleBeadResponse{bead}
}

func newBeadListResponse(c echo.Context, beads []model.Bead) *beadListResponse {
	r := new(beadListResponse)
	br := beadResponse{}
	r.Beads = make([]beadResponse, 0)
	for _, i := range beads {
		br.ID = i.ID
		br.Color = i.Color

		r.Beads = append(r.Beads, br)
	}
	return r
}
