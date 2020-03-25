package recs

import (
	"github.com/pocockn/recs-api/models"
)

// Store represents the database interactions.
type Store interface {
	Fetch(id uint) (rec models.Rec, err error)
	FetchAll() (recs models.Recs, err error)
	Update(rec *models.Rec) error
}
