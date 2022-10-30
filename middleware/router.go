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

	initializerGroup := r.Group("/initializer")
	migrationInitializerController := serviceContainer.InitMigrationInitializerController()
	initializerGroup.POST("/pokemonType", migrationInitializerController.InitPokemonTypes)
	initializerGroup.POST("/pokemon", migrationInitializerController.InitPokemons)

	pokedexGroup := r.Group("/pokedex")
	pokedexController := serviceContainer.InitPokedexController()
	pokedexGroup.GET("/:id", pokedexController.GetPokemon)
	pokedexGroup.GET("/byName/:name", pokedexController.GetPokemonByName)

	trainerGroup := r.Group("/trainer")
	trainerController := serviceContainer.InitTrainerController()
	trainerGroup.GET("/:id", trainerController.GetTrainer)
	trainerGroup.GET("/byNickname/:nickname", trainerController.GetTrainerByNickname)
	trainerGroup.POST("/", trainerController.CreateTrainer)

	return r
}

func NewRouter() IRouter {
	return &router{}
}
