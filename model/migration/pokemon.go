package migration

type Pokemon struct {
	Id          uint64         `json:"id" gorm:"primaryKey"`
	Name        PokemonName    `json:"name"`
	Species     string         `json:"species"`
	Description string         `json:"description"`
	Types       []string       `json:"type"`
	Image       PokemonImage   `json:"image"`
	Base        PokemonBase    `json:"base"`
	Profile     PokemonProfile `json:"profile"`
}
