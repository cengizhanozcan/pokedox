package service

import (
	"pokemon/model"
)

type IPokedexService interface {
	GetPokemon(id uint64) *model.Pokemon
	AddPokemon(pokemon *model.Pokemon)
}

type pokedexService struct {
	db *model.InMemoryDb
}

func (p *pokedexService) GetPokemon(id uint64) *model.Pokemon {
	for _, pokemon := range *p.db.GetPokemons() {
		if pokemon.Id == id {
			return &pokemon
		}
	}

	return nil
}

func (p *pokedexService) GetPokemonByName(name string) *model.Pokemon {
	for _, pokemon := range *p.db.GetPokemons() {
		if pokemon.Name.English == name {
			return &pokemon
		}
	}

	return nil
}

func (p *pokedexService) AddPokemon(pokemon *model.Pokemon) {
	if p.GetPokemonByName(pokemon.Name.English) != nil {
		return
	}

	p.db.AddPokemon(pokemon)
}

func NewPokedexService(db *model.InMemoryDb) IPokedexService {
	return &pokedexService{db: db}
}
