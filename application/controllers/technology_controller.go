package controllers

import (
	"net/http"

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

// CreateTechnologyHandler godoc
// @BasePath 	 	 /v1/api
// @Summary      Creates a new technology
// @Description  Receives a JSON with technology data and creates a new technology
// @Tags         technology
// @Accept       json
// @Produce      json
// @Param			  technology	body		models.TechnologyRequest	true	"Add technology"
// @Success     200  {object}  	usecases.TechnologyOutput
// @Failure		 	400		{object}	httputil.HTTPError
// @Failure		 	404		{object}	httputil.HTTPError
// @Failure		 	500		{object}	httputil.HTTPError
// @Router       			/technology [post]
func (h *TechnologyHandler) CreateTechnologyHandler(ctx *gin.Context) {
	request := models.TechnologyRequest{}
	ctx.BindJSON(&request)
	newTechnology, err := h.TechnologyUsecase.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(201, newTechnology)
}

// GetAllTechnologies godoc
// @BasePath 	 	 /v1/api
// @Summary      Gets all technologies
// @Description  returns all technologies
// @Tags         technology
// @Accept       json
// @Produce      json
// @Success      200  {object}  []usecases.TechnologyOutput
// @Failure		 	400		{object}	httputil.HTTPError
// @Failure		 	404		{object}	httputil.HTTPError
// @Failure		 	500		{object}	httputil.HTTPError
// @Router       			/technologies [get]
func (h *TechnologyHandler) GetAllTechnologies(ctx *gin.Context) {
	technologies, err := h.TechnologyUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, technologies)
}

