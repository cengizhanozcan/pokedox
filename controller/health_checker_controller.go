package controller

import "github.com/gin-gonic/gin"

type IHealthCheckerController interface {
	Ready(context *gin.Context)
}

type healthCheckerController struct {
}

func (h *healthCheckerController) Ready(context *gin.Context) {
	context.JSON(200, "Healthy")
}

func NewHealthChecker() IHealthCheckerController {
	return &healthCheckerController{}
}
