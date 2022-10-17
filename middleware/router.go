package middleware

import (
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	InitRouter() *gin.Engine
}

type router struct {
}

func (router *router) InitRouter() *gin.Engine {
	serviceContainer := NewServiceContainer()

	r := gin.New()
	healthGroup := r.Group("/_monitoring")
	healthGroup.GET("/ready", serviceContainer.InitHealthCheckerController().Ready)

	pokedexGroup := r.Group("/pokedex")
	pokedexController := serviceContainer.InitPokedexController()
	pokedexGroup.GET("/:id", pokedexController.GetPokemon)
	pokedexGroup.POST("/", pokedexController.AddPokemon)

	return r
}

func NewRouter() IRouter {
	return &router{}
}
