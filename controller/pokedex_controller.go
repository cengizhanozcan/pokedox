package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/service"
	"strconv"
)

type IPokedexController interface {
	GetPokemon(context *gin.Context)
	GetPokemonByName(context *gin.Context)
}

type pokedexController struct {
	pokedexService service.IPokedexService
}

func (p *pokedexController) GetPokemon(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
	pokemon := p.pokedexService.GetPokemon(id)
	context.JSON(http.StatusOK, pokemon)
}

func (p *pokedexController) GetPokemonByName(context *gin.Context) {
	name := context.Param("name")
	pokemon := p.pokedexService.GetPokemonByName(name)
	context.JSON(http.StatusOK, pokemon)
}

func NewPokedexController(pokedexService service.IPokedexService) IPokedexController {
	return &pokedexController{pokedexService: pokedexService}
}
