package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pokemon/service"
)

type IMigrationInitializerController interface {
	InitPokemonTypes(context *gin.Context)
	InitPokemons(context *gin.Context)
}

type migrationInitializerController struct {
	migrationInitializer service.IMigrationInitializer
}

func (p *migrationInitializerController) InitPokemonTypes(context *gin.Context) {
	p.migrationInitializer.InjectPokemonTypesData()
	context.Status(http.StatusOK)
	log.Println("Pokemon type initializer finished successfully")
}

func (p *migrationInitializerController) InitPokemons(context *gin.Context) {
	p.migrationInitializer.InjectPokemons()
	context.Status(http.StatusOK)
	log.Println("Pokemon initializer finished successfully")
}

func NewMigrationInitializerController(migrationInitializer service.IMigrationInitializer) IMigrationInitializerController {
	return &migrationInitializerController{migrationInitializer: migrationInitializer}
}
