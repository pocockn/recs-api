package delivery

import (
	"github.com/pocockn/recs-api/config"
	"github.com/pocockn/recs-api/models"
	"github.com/pocockn/recs-api/services"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/pocockn/recs-api/recs"
)

// Handler processes requests to our handlers from our API.
type Handler struct {
	Config        config.Config
	Store         recs.Store
	UploadService services.Upload
}

// NewHandler creates a new handler struct.
func NewHandler(config config.Config, store recs.Store) *Handler {
	return &Handler{
		Config:        config,
		Store:         store,
		UploadService: services.NewUpload(config, config.S3.Client),
	}
}

// Fetch gets a shout from it's ID.
func (h *Handler) Fetch(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	id := uint(idP)

	rec, err := h.Store.Fetch(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rec)
}

// FetchAll fetches all the recs from the DB.
func (h *Handler) FetchAll(c echo.Context) error {
	allRecs, err := h.Store.FetchAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, allRecs)
}

// Update updates a rec in the database.
func (h *Handler) Update(c echo.Context) error {
	rec := models.Rec{}
	err := c.Bind(&rec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = h.Store.Update(&rec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rec)
}
