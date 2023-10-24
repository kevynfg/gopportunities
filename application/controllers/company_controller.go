package controllers

import (
	"net/http"

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

// CreateCompanyHandler godoc
// @BasePath 	 	 /v1/api
// @Summary      Creates a new company
// @Description  Receives a JSON with company data and creates a new company
// @Tags         company
// @Accept       json
// @Produce      json
// @Param			  company	body		models.CompanyRequest	true	"Add company"
// @Success     200  {object}  	usecases.CompanyOutput
// @Failure		 	400		{object}	httputil.HTTPError
// @Failure		 	404		{object}	httputil.HTTPError
// @Failure		 	500		{object}	httputil.HTTPError
// @Router       			/company [post]
func (h *CompanyHandler) CreateCompanyHandler(ctx *gin.Context) {
	request := models.CompanyRequest{}
	ctx.BindJSON(&request)
	newTechnology, err := h.CompanyUsecase.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(201, newTechnology)
}

// GetAllCompanies godoc
// @BasePath 	 	 /v1/api
// @Summary      Gets all companies
// @Description  returns all companies
// @Tags         company
// @Accept       json
// @Produce      json
// @Success      200  {object}  []usecases.CompanyOutput
// @Failure		 	400		{object}	httputil.HTTPError
// @Failure		 	404		{object}	httputil.HTTPError
// @Failure		 	500		{object}	httputil.HTTPError
// @Router       			/companies [get]
func (h *CompanyHandler) GetAllCompanies(ctx *gin.Context) {
	companies, err := h.CompanyUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, companies)
}

