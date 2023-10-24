package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kevynfg/gopportunities/domain/models"
	"github.com/kevynfg/gopportunities/domain/usecases"
	"github.com/kevynfg/gopportunities/httputil"
)

type OpportunityHandler struct {
	OpportunityUsecase usecases.OpportunitiesUsecases
}

func NewOpportunityHandler(opportunityUsecase usecases.OpportunitiesUsecases) *OpportunityHandler {
	return &OpportunityHandler{OpportunityUsecase: opportunityUsecase}
}

// CreateOpportunityHandler godoc
// @BasePath 	 	 /v1/api
// @Summary      Creates a new opportunity
// @Description  Receives a JSON with opportunity data and creates a new opportunity
// @Tags         opportunity
// @Accept       json
// @Produce      json
// @Param			  opportunity	body		models.OpportunityRequest	true	"Add opportunity"
// @Success      200  {object}  models.OpportunityResponse
// @Failure		 	400		{object}	httputil.HTTPError
// @Failure		 	404		{object}	httputil.HTTPError
// @Failure		 	500		{object}	httputil.HTTPError
// @Router       			/opportunity [post]
func (h *OpportunityHandler) CreateOpportunityHandler(ctx *gin.Context) {
	request := models.OpportunityRequest{}
	ctx.BindJSON(&request)
	newOpportunity, err := h.OpportunityUsecase.CreateOpportunity(request)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(201, &newOpportunity)
}


// GetAllOpportunities godoc
// @BasePath 	 	 /v1/api
// @Summary      Gets all opportunities
// @Description  Receives a query with limit and offset and returns all opportunities
// @Tags         opportunity
// @Accept       json
// @Produce      json
// @Param   limit    query    int     false        "Limit"
// @Param   offset   query    int     false        "Offset"
// @Success      200  {object}  []models.OpportunityResponse
// @Failure		 	400		{object}	httputil.HTTPError
// @Failure		 	404		{object}	httputil.HTTPError
// @Failure		 	500		{object}	httputil.HTTPError
// @Router       			/opportunities [get]
func (h *OpportunityHandler) GetAllOpportunities(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	opportunities, err := h.OpportunityUsecase.FindAll(limit, offset)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, opportunities)
}

// UpdateOpportunityHandler godoc
// @BasePath 	 	 /v1/api
// @Summary      Updates a opportunity
// @Description  Receives a JSON with opportunity data and creates a new opportunity
// @Tags         opportunity
// @Accept       json
// @Produce      json
// @Param				 id		path		int					true	"Opportunity ID"
// @Param			   opportunity	body		models.OpportunityRequest	true	"Update opportunity"
// @Success      200    {object}  models.OpportunityResponse
// @Failure		 	 400		{object}	httputil.HTTPError
// @Failure		 	 404		{object}	httputil.HTTPError
// @Failure		 	 500		{object}	httputil.HTTPError
// @Router       			  /opportunity/:id [put]
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
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, opportunity)
}

func DisableOpportunityHandler(ctx *gin.Context) {
		
}