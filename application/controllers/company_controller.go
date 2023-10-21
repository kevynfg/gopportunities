package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/domain/usecases"
)

type CompanyHandler struct {
	CompanyUsecase usecases.CompaniesUsecases
}

func NewCompanyHandler(companyUsecase usecases.CompaniesUsecases) *CompanyHandler {
	return &CompanyHandler{CompanyUsecase: companyUsecase}
}

func (h *CompanyHandler) CreateCompanyHandler(ctx *gin.Context) {
	request := models.CompanyRequest{}
	ctx.BindJSON(&request)
	newTechnology, err := h.CompanyUsecase.Execute(request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, newTechnology)
}

func (h *CompanyHandler) GetAllCompanies(ctx *gin.Context) {
	companies, err := h.CompanyUsecase.FindAll()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, companies)
}

