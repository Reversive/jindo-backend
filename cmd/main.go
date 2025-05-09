package main

import (
	"log"

	"github.com/Reversive/jindo/internal/handlers"
	"github.com/Reversive/jindo/internal/repositories"
	"github.com/Reversive/jindo/internal/router"
	"github.com/Reversive/jindo/internal/services"
	"github.com/Reversive/jindo/internal/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config: ", err)
	}

	r := router.NewRouter()
	novelRepository := repositories.NewNovelRepository(config.DBSource)
	defer novelRepository.Close()
	novelService := services.NewNovelService(novelRepository)
	novelHandler := handlers.NewNovelHandler(novelService)

	r.Post("/api/novels", novelHandler.PostNovel)

	r.Run(config.ServerAdress)
	return
}
