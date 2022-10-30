package domain

import (
	"gorm.io/gorm"
	"time"
)

type Trainer struct {
	gorm.Model
	ID        uint64
	Name      string
	Surname   string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
