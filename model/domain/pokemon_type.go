package domain

import (
	"gorm.io/gorm"
	"time"
)

type PokemonType struct {
	gorm.Model
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
