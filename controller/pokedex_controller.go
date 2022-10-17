package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/model"
	"pokemon/service"
	"strconv"
)

type IPokedexController interface {
	GetPokemon(context *gin.Context)
	AddPokemon(context *gin.Context)
}

type pokedexController struct {
	pokedexService service.IPokedexService
}

func (p *pokedexController) GetPokemon(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
	pokemon := p.pokedexService.GetPokemon(id)
	context.JSON(http.StatusOK, pokemon)
}

func (p *pokedexController) AddPokemon(context *gin.Context) {
	var pokemon model.Pokemon
	if err := context.BindJSON(&pokemon); err != nil {
		return
	}
	p.pokedexService.AddPokemon(&pokemon)
	context.Status(http.StatusOK)
}

func NewPokedexController(pokedexService service.IPokedexService) IPokedexController {
	return &pokedexController{pokedexService: pokedexService}
}
