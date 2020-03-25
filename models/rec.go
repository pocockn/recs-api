package models

import "github.com/jinzhu/gorm"

type (
	// Rec holds all the information related to a recommendation.
	Rec struct {
		gorm.Model
		Rating    int64  `json:"rating"`
		Review    string `json:"review"`
		SpotifyID string `json:"spotify_id"`
		Title     string `json:"title"`
	}

	// Recs holds a list of Rec structs.
	Recs []Rec
)
