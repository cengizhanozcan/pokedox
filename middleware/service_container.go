package middleware

import (
	"pokemon/controller"
	"pokemon/model"
	"pokemon/service"
)

type IServiceContainer interface {
	InitHealthCheckerController() controller.IHealthCheckerController
	InitPokedexController() controller.IPokedexController
}

type serviceContainer struct {
	db *model.InMemoryDb
}

func (s *serviceContainer) InitHealthCheckerController() controller.IHealthCheckerController {
	return controller.NewHealthChecker()
}

func (s *serviceContainer) InitPokedexController() controller.IPokedexController {
	pokedexService := service.NewPokedexService(s.db)
	return controller.NewPokedexController(pokedexService)
}

func NewServiceContainer() IServiceContainer {
	db := model.NewInMemoryDb(GetDataFromJson())
	return &serviceContainer{db: db}
}
