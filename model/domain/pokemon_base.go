package domain

import (
	"gorm.io/gorm"
	"time"
)

type PokemonBase struct {
	gorm.Model
	PokemonID uint64
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
