package model

type InMemoryDb struct {
	pokemons *[]Pokemon
}

func (i *InMemoryDb) GetPokemons() *[]Pokemon {
	return i.pokemons
}

func (i *InMemoryDb) AddPokemon(pokemon *Pokemon) {
	*i.pokemons = append(*i.pokemons, *pokemon)
}

func NewInMemoryDb(pokemons *[]Pokemon) *InMemoryDb {
	return &InMemoryDb{pokemons: pokemons}
}
