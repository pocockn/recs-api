package handler

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

// RecHandler implements the store interface.
type RecHandler struct {
	Config        config.Config
	Repo          recs.Store
	UploadService services.Upload
}

// NewRecHandler creates a new recs handler with the routes.
func NewRecHandler(config config.Config, repo recs.Store) *RecHandler {
	return &RecHandler{
		Config:        config,
		Repo:          repo,
		UploadService: services.NewUpload(config, config.S3.Client),
	}
}

// Fetch gets a shout from it's ID.
func (s *RecHandler) Fetch(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	id := uint(idP)

	rec, err := s.Repo.Fetch(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rec)
}

// FetchAll fetches all the recs from the DB.
func (s *RecHandler) FetchAll(c echo.Context) error {
	allRecs, err := s.Repo.FetchAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, allRecs)
}

// Update updates a rec in the database.
func (s *RecHandler) Update(c echo.Context) error {
	rec := models.Rec{}
	err := c.Bind(&rec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = s.Repo.Update(&rec)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, rec)
}

//// Store takes a shout and stores it in the DB.
//func (s *RecHandler) Store(c echo.Context) error {
//	guid := uuid.NewV4().String()
//	sourceFile, err := c.FormFile("source")
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, nil)
//	}
//
//	targetFile, err := c.FormFile("target")
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, nil)
//	}
//
//	err = s.UploadService.MultiUpload(sourceFile, targetFile)
//	if err != nil {
//		return err
//	}
//
//	imageSimilarityPayload, err := lambda.NewRawJsonImageSimilarity(
//		targetFile.Filename,
//		sourceFile.Filename,
//		guid,
//	)
//	if err != nil {
//		return err
//	}
//
//	snsMessage := sns.Message{
//		ID:      guid,
//		Payload: &imageSimilarityPayload,
//	}
//
//	snsMessagePayload, err := json.Marshal(&snsMessage)
//	if err != nil {
//		return err
//	}
//
//	snsClient := *awsWrappersSNS.NewClient(nil, false, nil)
//	messageID, err := snsClient.PublishMessage(
//		string(snsMessagePayload),
//		s.Config.SNS.Arn,
//	)
//	if err != nil {
//		return err
//	}
//
//	log.Infof("SNS notification %s", messageID)
//
//	shout := models.
//
//	s.Repo.Store()
//
//	return c.JSON(http.StatusCreated, nil)
//}
