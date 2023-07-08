package models

// import "gorm.io/gorm"

// album represents data about a record album.
type Album struct {
	// gorm.Model
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
