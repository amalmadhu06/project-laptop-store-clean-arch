package main

import (
	_ "github.com/amalmadhu06/project-laptop-store-clean-arch/cmd/api/docs"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/di"
	"log"
)

// @title Ecommerce REST API
// @version 1.0
// @description Ecommerce REST API built using Go Lang, PSQL, REST API following Clean Architecture. Hosted with Ngnix, AWS EC2 and RDS

// @contact.name Amal Madhu
// @contact.url https://github.com/amalmadhu06
// @contact.email madhuamal06@gmail.com

// @license.name MIT
// @host amalmadhu.live
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
