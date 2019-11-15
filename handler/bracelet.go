package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang-starter-pack/model"
	"golang-starter-pack/utils"
)

func (h *Handler) GetBracelet(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.braceletStore.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newBraceletResponse(c, a))
}

func (h *Handler) Bracelets(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	var bracelets []model.Bracelet
	
	return c.JSON(http.StatusOK, newBraceletListResponse(bracelets))
}

func (h *Handler) CreateBracelet(c echo.Context) error {
	var a model.Bracelet
	req := &braceletCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err := h.braceletStore.CreateBracelet(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newBraceletResponse(c, &a))
}

func (h *Handler) DeleteBracelet(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.braceletStore.GetUserBraceletBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = h.braceletStore.DeleteBracelet(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

func (h *Handler) AddBead(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.braceletStore.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	var cm model.Bead
	req := &createBeadRequest{}
	if err := req.bind(c, &cm); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = h.braceletStore.AddBead(a, &cm); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newBeadResponse(c, &cm))
}

func (h *Handler) GetBeads(c echo.Context) error {
	slug := c.Param("slug")
	cm, err := h.braceletStore.GetBeadsBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newBeadListResponse(c, cm))
}

func (h *Handler) DeleteBead(c echo.Context) error {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}
	cm, err := h.braceletStore.GetBeadByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if cm == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	if err := h.braceletStore.DeleteBead(cm); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

