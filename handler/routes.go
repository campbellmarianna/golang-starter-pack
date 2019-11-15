package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	bracelets := v1.Group("/bracelets")
	bracelets.POST("", h.CreateBracelet)
	bracelets.DELETE("/:slug", h.DeleteBracelet)
	bracelets.POST("/:slug/beads", h.AddBead)
	bracelets.DELETE("/:slug/beads/:id", h.DeleteBead)
	bracelets.GET("/:slug", h.GetBracelet)
	bracelets.GET("/:slug/beads", h.GetBeads)
}
