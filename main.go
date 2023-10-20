package main

import (
	db "github.com/kevynfg/gopportunities/infra"
	"github.com/kevynfg/gopportunities/router"
)

func main() {
  db.InitDB()
  router.Initialize()
}