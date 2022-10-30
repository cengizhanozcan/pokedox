package migration

type PokemonBase struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SpAttack  int `json:"Sp. Attack"`
	SpDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}
