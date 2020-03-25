package handler

import "github.com/labstack/echo"

// Registers routes for our echo api.
func (h *RecHandler) Register(e *echo.Echo) {
	e.GET("/recs/:id", h.Fetch)
	e.GET("/recs", h.FetchAll)
	e.POST("/rec", h.Update)
}
