package migration

type PokemonImage struct {
	Sprite    string `json:"sprite"`
	Thumbnail string `json:"thumbnail"`
	Hires     string `json:"hires"`
}
