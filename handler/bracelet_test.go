package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang-starter-pack/router"
	"golang-starter-pack/router/middleware"
	"golang-starter-pack/utils"
)

func TestListBraceletsCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/bracelets", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, h.Bracelets(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var aa braceletListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &aa)
		assert.NoError(t, err)
		assert.Equal(t, 2, aa.BraceletsCount)
	}
}

func TestGetBraceletsCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	req := httptest.NewRequest(echo.GET, "/api/bracelets/:slug", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/bracelets/:slug")
	c.SetParamNames("slug")
	c.SetParamValues("bracelet1-slug")
	assert.NoError(t, h.GetBracelet(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var a singleBraceletResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "bracelet1-slug", a.Bracelet.Slug)
	}
}

func TestCreateBraceletsCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"bracelet":{"text":"Black Girls Code", "threadColor":"blue", "font":"bold playful"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/bracelets", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Might not work without jwtMiddleware wrapper
	err := func(context echo.Context) error {
		return h.CreateBracelet(c)
	}(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var a singleBraceletResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "Black-Girls-Code", a.Bracelet.Slug)
		assert.Equal(t, "blue", a.Bracelet.ThreadColor)
		assert.Equal(t, "Black Girls Code", a.Bracelet.Text)
	}
}
// I might this in when I want to have an Update function
// func TestUpdateBraceletsCaseSuccess(t *testing.T) {
// 	tearDown()
// 	setup()
// 	var (
// 		reqJSON = `{"bracelet":{"text":"Naturally MEE", "threadColor":"blue", "font":"curvy"}}`
// 	)
// 	req := httptest.NewRequest(echo.PUT, "/api/bracelets/:slug", strings.NewReader(reqJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/api/bracelets/:slug")
// 	c.SetParamNames("slug")
// 	c.SetParamValues("Naturally-MEE-slug")
// 	err := func(context echo.Context) error {
// 		return h.UpdateBracelet(c)
// 	}(c)
// 	assert.NoError(t, err)
// 	if assert.Equal(t, http.StatusOK, rec.Code) {
// 		var a singleBraceletResponse
// 		err := json.Unmarshal(rec.Body.Bytes(), &a)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Naturally-MEE", a.Bracelet.Text)
// 		assert.Equal(t, "Naturally-MEE-Naturally-MEE", a.Bracelet.Slug)
// 	}
// }

func TestDeleteBraceletCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	req := httptest.NewRequest(echo.DELETE, "/api/bracelets/:slug", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/bracelets/:slug")
	c.SetParamNames("slug")
	c.SetParamValues("Black-Girls-Code-slug")
	err := func(context echo.Context) error {
		return h.DeleteBracelet(c)
	}(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBeadsCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	req := httptest.NewRequest(echo.GET, "/api/bracelets/:slug/beads", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/bracelets/:slug/beads")
	c.SetParamNames("slug")
	c.SetParamValues("Black-Girls-Code-slug")
	err := func(context echo.Context) error {
		return h.GetBeads(c)
	}(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var cc beadListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &cc)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(cc.Beads))
	}
}

func TestAddBeadCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"bead":{"color":"pink", "design": "square pattern"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/bracelets/:slug/comments", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/bracelets/:slug/beads")
	c.SetParamNames("slug")
	c.SetParamValues("Black-Girls-Code-slug")
	err := func(context echo.Context) error {
		return h.AddBead(c)
	}(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var c singleBeadResponse
		err := json.Unmarshal(rec.Body.Bytes(), &c)
		assert.NoError(t, err)
		assert.Equal(t, "pink", c.Bead.Color)
	}
}

func TestDeleteBeadCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	req := httptest.NewRequest(echo.DELETE, "/api/bracelets/:slug/beads/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/bracelets/:slug/beads/:id")
	c.SetParamNames("slug")
	c.SetParamValues("Black-Girls-Code-slug")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := func(context echo.Context) error {
		return h.DeleteComment(c)
	}(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

