package model

type Pokemon struct {
	Id    uint64       `json:"id"`
	Name  PokemonName  `json:"name"`
	Types []string     `json:"type"`
	Image PokemonImage `json:"image"`
}
