package recs

import (
	"github.com/pocockn/recs-api/models"
)

// Store represents the rec's store context.
type Store interface {
	Fetch(id uint) (rec models.Rec, err error)
	FetchAll() (recs models.Recs, err error)
	Update(rec *models.Rec) error
}
