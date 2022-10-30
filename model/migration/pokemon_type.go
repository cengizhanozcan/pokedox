package migration

type PokemonType struct {
	English     string        `json:"english"`
	Chinese     string        `json:"chinese"`
	Japanese    string        `json:"japanese"`
	Effective   []interface{} `json:"effective"`
	Ineffective []string      `json:"ineffective"`
	NoEffect    []string      `json:"no_effect"`
}
