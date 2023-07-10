package models

import (
	"time"
)

type Model struct {
	CreatedAt time.Time `gorm:"<-:false;type:datetime NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"<-:false;type:datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
