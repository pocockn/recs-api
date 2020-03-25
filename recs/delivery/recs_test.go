package delivery_test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pocockn/recs-api/config"
	"github.com/pocockn/recs-api/models"
	"github.com/pocockn/recs-api/recs/delivery"
	mock_recs "github.com/pocockn/recs-api/recs/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_recs.NewMockStore(controller)

	mockStore.EXPECT().
		Fetch(uint(1)).
		Return(models.Rec{
			Model:     gorm.Model{ID: 1},
			Rating:    2,
			Review:    "Meh",
			SpotifyID: "1234",
			Title:     "Hello",
		}, nil)

	e := echo.New()
	h := delivery.NewHandler(config.Config{}, mockStore)

	req := httptest.NewRequest(echo.GET, "/recs/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/recs/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.Fetch(c))
}

func TestFetchAll(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_recs.NewMockStore(controller)

	mockStore.EXPECT().
		FetchAll().
		Return(
			models.Recs{
				models.Rec{
					Model:     gorm.Model{ID: 1},
					Rating:    2,
					Review:    "Meh",
					SpotifyID: "1234",
					Title:     "Hello",
				},
			}, nil)

	e := echo.New()
	h := delivery.NewHandler(config.Config{}, mockStore)

	req := httptest.NewRequest(echo.GET, "/recs", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.FetchAll(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var recs models.Recs
		err := json.Unmarshal(rec.Body.Bytes(), &recs)
		assert.NoError(t, err)
		assert.Equal(t, "Hello", recs[0].Title)
		assert.Equal(t, 1, len(recs))
	}
}

func TestFetchError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_recs.NewMockStore(controller)
	dummyErr := errors.New("fetch all error")

	mockStore.EXPECT().
		FetchAll().
		Return(nil, dummyErr)

	e := echo.New()
	h := delivery.NewHandler(config.Config{}, mockStore)

	req := httptest.NewRequest(echo.GET, "/recs", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.FetchAll(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestUpdate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_recs.NewMockStore(controller)
	updateRec := models.Rec{
		Rating: 2,
	}
	reqJSON := `{"rating":2}`

	mockStore.EXPECT().
		Update(&updateRec).
		Return(nil)

	e := echo.New()
	h := delivery.NewHandler(config.Config{}, mockStore)

	req := httptest.NewRequest(echo.POST, "/rec", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.Update(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var returnedRec models.Rec
		err := json.Unmarshal(rec.Body.Bytes(), &returnedRec)
		assert.NoError(t, err)
		assert.Equal(t, 2, int(returnedRec.Rating))
	}
}

func TestUpdateError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mock_recs.NewMockStore(controller)
	updateRec := models.Rec{
		Rating: 2,
	}
	reqJSON := `{"rating":2}`

	mockStore.EXPECT().
		Update(&updateRec).
		Return(errors.New("update error"))

	e := echo.New()
	h := delivery.NewHandler(config.Config{}, mockStore)

	req := httptest.NewRequest(echo.POST, "/rec", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.Update(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
