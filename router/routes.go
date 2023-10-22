package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kevynfg/gopportunities/application/controllers"
	"github.com/kevynfg/gopportunities/factories"
	db "github.com/kevynfg/gopportunities/infra"
	"github.com/kevynfg/gopportunities/infra/repositories"
)

var dbConnection = db.GetDB()

func InitRoutes(router *gin.Engine) {
	technologyRepository := repositories.NewTechnologiesRepositorySql(dbConnection)
	technologyHandler := factories.NewTechnologyController(*technologyRepository)
	companyRepository := repositories.NewCompaniesRepositorySql(dbConnection)
	companyHandler := factories.NewCompanyController(*companyRepository)
	opportunityRepository := repositories.NewOpportunitiesRepositorySql(dbConnection)
	opportunityHandler := factories.NewOpportunityController(*opportunityRepository)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", opportunityHandler.GetAllOpportunities)
		v1.GET("/technologies", technologyHandler.GetAllTechnologies)
		v1.GET("/companies", companyHandler.GetAllCompanies)
		v1.POST("/opportunity", opportunityHandler.CreateOpportunityHandler)
		v1.POST("/technology", technologyHandler.CreateTechnologyHandler)
		v1.POST("/company", companyHandler.CreateCompanyHandler)
		v1.PUT("/opportunity/:id", opportunityHandler.UpdateOpportunityHandler)
		v1.GET("/opportunities/search", opportunityHandler.GetAllOpportunities)
		v1.DELETE("/opportunity/:id", controllers.DisableOpportunityHandler)
	}
	router.Run()
}
