package store

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/recs-api/models"
	"github.com/pocockn/recs-api/recs"
)

type recsStore struct {
	Conn *gorm.DB
}

// NewRecsStore creates a new recsStore struct for interacting with the Gorm connection to the DB.
func NewRecsStore(conn *gorm.DB) recs.Store {
	return &recsStore{conn}
}

// Fetch fetches a shout via it's ID from the DB.
func (s *recsStore) Fetch(id uint) (models.Rec, error) {
	var rec models.Rec
	err := s.Conn.Where("id = ?", id).First(&rec).Error

	return rec, err
}

// FetchAll fetches all the recs from the DB.
func (s *recsStore) FetchAll() (models.Recs, error) {
	var recs models.Recs
	err := s.Conn.Find(&recs).Error

	return recs, err
}

// Update updates a rec in the DB.
func (s *recsStore) Update(rec *models.Rec) error {
	return s.Conn.Save(&rec).Error
}
