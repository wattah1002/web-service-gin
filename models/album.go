package models

import "gorm.io/gorm"

// album represents data about a record album.
type Album struct {
	ID        string  `gorm:"primaryKey;type:char(36);not null"`
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Price     float64 `json:"price"`
	DeletedAt gorm.DeletedAt
	Model
}
