package service

import (
	"pokemon/middleware/database"
	"pokemon/model/domain"
)

type IPokedexService interface {
	GetPokemon(id uint64) *domain.Pokemon
	GetPokemonByName(name string) *domain.Pokemon
}

type pokedexService struct {
	db *database.Database
}

func (p *pokedexService) GetPokemon(id uint64) *domain.Pokemon {
	pokemon := domain.Pokemon{}
	p.db.Instance.
		Preload("Types").
		Preload("PokemonImage").
		Preload("PokemonBase").
		Preload("PokemonProfile").
		First(&pokemon, "id=?", id)
	if pokemon.ID != 0 {
		return &pokemon
	}
	return nil
}

func (p *pokedexService) GetPokemonByName(name string) *domain.Pokemon {
	pokemon := domain.Pokemon{}
	p.db.Instance.First(&pokemon, "name=?", name)
	if pokemon.ID != 0 {
		return &pokemon
	}
	return nil
}

func NewPokedexService(db *database.Database) IPokedexService {
	return &pokedexService{db: db}
}
