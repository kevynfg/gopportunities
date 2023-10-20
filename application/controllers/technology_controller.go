package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/domain/usecases"
)

type TechnologyHandler struct {
	TechnologyUsecase usecases.TechnologiesUsecases
}

func NewTechnologyHandler(technologyUsecase usecases.TechnologiesUsecases) *TechnologyHandler {
	return &TechnologyHandler{TechnologyUsecase: technologyUsecase}
}

func (h *TechnologyHandler) CreateTechnologyHandler(ctx *gin.Context) {
	request := models.TechnologyRequest{}
	ctx.BindJSON(&request)
	newTechnology, err := h.TechnologyUsecase.Execute(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, newTechnology)
}

func (h *TechnologyHandler) GetAllTechnologies(ctx *gin.Context) {
	technologies, err := h.TechnologyUsecase.FindAll()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, technologies)
}

