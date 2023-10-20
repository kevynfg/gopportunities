package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynfg/gopportunities/application/controllers"
	"github.com/kevynfg/gopportunities/domain/usecases"
	db "github.com/kevynfg/gopportunities/infra"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

var dbConnection = db.GetDB()

func InitRoutes(router *gin.Engine) {
	repository := repositories.NewTechnologiesRepositorySql(dbConnection)
	technologyUsecase:= usecases.NewTechnologyUsecases(*repository)
	technologyHandler := controllers.NewTechnologyHandler(*technologyUsecase)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", controllers.GetOpportunitiesHandler)
		v1.GET("/technologies", technologyHandler.GetAllTechnologies)
		v1.POST("/opportunity", controllers.CreateOpportunityHandler)
		v1.POST("/technology", technologyHandler.CreateTechnologyHandler)
		v1.PUT("/opportunity/:id", controllers.UpdateOpportunityHandler)
		v1.GET("/opportunities/search", controllers.SearchOpportunitiesHandler)
		v1.DELETE("/opportunity/:id", controllers.DisableOpportunityHandler)
	}
	router.Run()
}
