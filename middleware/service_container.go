package middleware

import (
	"pokemon/controller"
	"pokemon/middleware/database"
	"pokemon/service"
)

type IServiceContainer interface {
	InitHealthCheckerController() controller.IHealthCheckerController
	InitPokedexController() controller.IPokedexController
	InitMigrationInitializerController() controller.IMigrationInitializerController
	InitTrainerController() controller.ITrainerController
}

type serviceContainer struct {
	db *database.Database
}

func (s *serviceContainer) InitHealthCheckerController() controller.IHealthCheckerController {
	return controller.NewHealthChecker()
}

func (s *serviceContainer) InitPokedexController() controller.IPokedexController {
	pokedexService := service.NewPokedexService(s.db)
	return controller.NewPokedexController(pokedexService)
}

func (s *serviceContainer) InitMigrationInitializerController() controller.IMigrationInitializerController {
	initializer := service.NewMigrationInitializer(s.db)
	return controller.NewMigrationInitializerController(initializer)
}

func (s *serviceContainer) InitTrainerController() controller.ITrainerController {
	trainerService := service.NewTrainerService(s.db)
	return controller.NewTrainerController(trainerService)
}

func NewServiceContainer() IServiceContainer {
	db := database.NewInMemoryDb()
	return &serviceContainer{db: db}
}
