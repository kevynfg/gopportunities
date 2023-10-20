package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpportunityHandler(ctx *gin.Context) {
	
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
	//get all params from query
	ctx.JSON(http.StatusOK, gin.H{"query": query})
}