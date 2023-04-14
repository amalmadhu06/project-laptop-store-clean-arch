package main

import (
	_ "github.com/amalmadhu06/project-laptop-store-clean-arch/cmd/api/docs"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/di"
	"log"
)

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
