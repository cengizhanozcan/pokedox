package service

import (
	"encoding/json"
	"io/ioutil"
	"pokemon/middleware/database"
	"pokemon/model/domain"
	"pokemon/model/migration"
	"time"
)

type IMigrationInitializer interface {
	InjectPokemonTypesData()
	InjectPokemons()
}

type MigrationInitializer struct {
	db *database.Database
}

func (m *MigrationInitializer) InjectPokemonTypesData() {
	pokemonTypes := getPokemonTypeDataFromJson()
	for _, tempPokemonType := range *pokemonTypes {
		pokemonType := getPokemonType(tempPokemonType)
		m.db.Instance.Create(pokemonType)
	}
}

func (m *MigrationInitializer) InjectPokemons() {
	pokemons := getPokemonFromDataJson()
	for _, tempPokemon := range *pokemons {
		pokemon := m.getPokemon(tempPokemon)
		m.db.Instance.Create(pokemon)
	}
}

func getPokemonTypeDataFromJson() *[]migration.PokemonType {
	file, _ := ioutil.ReadFile("types.json")
	data := make([]migration.PokemonType, 1)
	_ = json.Unmarshal(file, &data)
	return &data
}

func getPokemonType(pokemonType migration.PokemonType) *domain.PokemonType {
	return &domain.PokemonType{
		Name:      pokemonType.English,
		CreatedAt: time.Time{},
	}
}

func getPokemonFromDataJson() *[]migration.Pokemon {
	file, _ := ioutil.ReadFile("pokedex.json")
	data := make([]migration.Pokemon, 1)
	_ = json.Unmarshal(file, &data)
	return &data
}

func (m *MigrationInitializer) getPokemon(pokemon migration.Pokemon) *domain.Pokemon {
	pokemonTypes := make([]domain.PokemonType, 0)
	for _, tempPokemonType := range pokemon.Types {
		pokemonType := domain.PokemonType{}
		m.db.Instance.First(&pokemonType, "name=?", tempPokemonType)
		pokemonTypes = append(pokemonTypes, pokemonType)
	}
	return &domain.Pokemon{
		ID:    pokemon.Id,
		Name:  pokemon.Name.English,
		Types: pokemonTypes,
		PokemonImage: &domain.PokemonImage{
			PokemonID: pokemon.Id,
			Sprite:    pokemon.Image.Sprite,
			Thumbnail: pokemon.Image.Thumbnail,
			Hires:     pokemon.Image.Hires,
			CreatedAt: time.Time{},
		},
		PokemonBase: &domain.PokemonBase{
			PokemonID: pokemon.Id,
			HP:        pokemon.Base.HP,
			Attack:    pokemon.Base.Attack,
			Defense:   pokemon.Base.Defense,
			SpAttack:  pokemon.Base.SpAttack,
			SpDefense: pokemon.Base.SpDefense,
			Speed:     pokemon.Base.Speed,
			CreatedAt: time.Time{},
			UpdatedAt: nil,
			DeletedAt: nil,
		},
		PokemonProfile: &domain.PokemonProfile{
			PokemonID: pokemon.Id,
			Height:    pokemon.Profile.Height,
			Weight:    pokemon.Profile.Weight,
			Gender:    pokemon.Profile.Height,
			CreatedAt: time.Time{},
			UpdatedAt: nil,
			DeletedAt: nil,
		},
		CreatedAt: time.Time{},
	}
}

func NewMigrationInitializer(db *database.Database) *MigrationInitializer {
	return &MigrationInitializer{db: db}
}
