package domain

import (
	"gorm.io/gorm"
	"time"
)

type PokemonImage struct {
	gorm.Model
	PokemonID uint64
	Sprite    string
	Thumbnail string
	Hires     string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
