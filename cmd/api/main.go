package main

import (
	_ "github.com/amalmadhu06/project-laptop-store-clean-arch/cmd/api/docs"

	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/di"
	"log"
)

// @title Ecommerce Web API
// @version 1.0
// @description Ecommerce Web Application built using Go Lang, PSQL, REST API following Clean Architecture.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @host localhost:3000
// @license.url https://opensource.org/licenses/MIT

// @BasePath /
// @query.collection.format multi
func main() {
	cfg, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load cfg: ", configErr)
	}

	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
