package main

import (
	"github.com/kevynfg/gopportunities/docs"
	db "github.com/kevynfg/gopportunities/infra"
	"github.com/kevynfg/gopportunities/router"
)

func main() {
  db.InitDB()
	docs.SwaggerInfo.Title = "Gopportunities API"
	docs.SwaggerInfo.Description = "A simple api to get job opportunities"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
  router.Initialize()
}