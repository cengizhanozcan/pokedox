package domain

import (
	"gorm.io/gorm"
	"time"
)

type PokemonProfile struct {
	gorm.Model
	PokemonID uint64
	Height    string
	Weight    string
	Gender    string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
