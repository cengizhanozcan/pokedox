package service

import (
	"pokemon/middleware/database"
	"pokemon/model/migration"
)

type IPokedexService interface {
	GetPokemon(id uint64) *migration.Pokemon
	AddPokemon(pokemon *migration.Pokemon)
}

type pokedexService struct {
	db *database.Database
}

func (p *pokedexService) GetPokemon(id uint64) *migration.Pokemon {
	for _, pokemon := range *p.db.GetPokemons() {
		if pokemon.ID == id {
			//return &pokemon
		}
	}

	return nil
}

func (p *pokedexService) GetPokemonByName(name string) *migration.Pokemon {
	for _, pokemon := range *p.db.GetPokemons() {
		if pokemon.Name == name {
			//return &pokemon
		}
	}

	return nil
}

func (p *pokedexService) AddPokemon(pokemon *migration.Pokemon) {
	if p.GetPokemonByName(pokemon.Name.English) != nil {
		return
	}

	//p.db.AddPokemon(pokemon)
}

func NewPokedexService(db *database.Database) IPokedexService {
	return &pokedexService{db: db}
}
