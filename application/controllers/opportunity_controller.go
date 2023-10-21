package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/domain/usecases"
)

type OpportunityHandler struct {
	OpportunityUsecase usecases.OpportunitiesUsecases
}

func NewOpportunityHandler(opportunityUsecase usecases.OpportunitiesUsecases) *OpportunityHandler {
	return &OpportunityHandler{OpportunityUsecase: opportunityUsecase}
}

func (h *OpportunityHandler) CreateOpportunityHandler(ctx *gin.Context) {
	request := models.OpportunityRequest{}
	ctx.BindJSON(&request)
	newOpportunity, err := h.OpportunityUsecase.CreateOpportunity(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, newOpportunity)
}

func GetOpportunitiesHandler(ctx *gin.Context) {
	
}

func UpdateOpportunityHandler(ctx *gin.Context) {
	ctx.Param("id")
	if ctx.Param("id") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": ctx.Param("id")})
}

func DisableOpportunityHandler(ctx *gin.Context) {
		
}

func SearchOpportunitiesHandler(ctx *gin.Context) {
	query := ctx.Query("name")
	ctx.JSON(http.StatusOK, gin.H{"query": query})
}