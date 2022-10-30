package domain

import (
	"gorm.io/gorm"
	"time"
)

type PokemonBase struct {
	gorm.Model
	PokemonID uint64
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
