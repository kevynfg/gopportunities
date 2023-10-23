package controllers

import (
	"net/http"
	"strconv"

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
	ctx.JSON(201, &newOpportunity)
}

func (h *OpportunityHandler) GetAllOpportunities(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	opportunities, err := h.OpportunityUsecase.FindAll(limit, offset)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, opportunities)
}

func (h *OpportunityHandler) UpdateOpportunityHandler(ctx *gin.Context) {
	ctx.Param("id")
	if ctx.Param("id") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	request := models.Opportunity{}
	ctx.BindJSON(&request)
	paramId, errAtoi := strconv.Atoi(ctx.Param("id"))
	if errAtoi != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	request.ID = uint(paramId)
	opportunity, err := h.OpportunityUsecase.EditOpportunity(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, opportunity)
}

func DisableOpportunityHandler(ctx *gin.Context) {
		
}