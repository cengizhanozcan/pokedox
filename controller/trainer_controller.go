package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pokemon/model/domain"
	"pokemon/service"
	"strconv"
)

type ITrainerController interface {
	GetTrainer(context *gin.Context)
	GetTrainerByNickname(context *gin.Context)
	CreateTrainer(context *gin.Context)
}

type trainerController struct {
	trainerService service.ITrainerService
}

func (t *trainerController) GetTrainer(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 64)
	trainer := t.trainerService.GetTrainer(id)
	context.JSON(http.StatusOK, trainer)
}

func (t *trainerController) GetTrainerByNickname(context *gin.Context) {
	nickname := context.Param("nickname")
	trainer := t.trainerService.GetTrainerByNickname(nickname)
	context.JSON(http.StatusOK, trainer)
}

func (t *trainerController) CreateTrainer(context *gin.Context) {
	var trainer domain.Trainer
	if err := context.BindJSON(&trainer); err != nil {
		return
	}
	t.trainerService.CreateTrainer(&trainer)
	context.Status(http.StatusOK)
}

func NewTrainerController(trainerService service.ITrainerService) ITrainerController {
	return &trainerController{trainerService: trainerService}
}
