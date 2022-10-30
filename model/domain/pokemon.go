package domain

import (
	"gorm.io/gorm"
	"time"
)

type Pokemon struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey"`
	Name           string
	Types          []PokemonType `gorm:"many2many:pokemon_pokemon_types;"`
	PokemonImage   *PokemonImage
	PokemonBase    *PokemonBase
	PokemonProfile *PokemonProfile
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
}
