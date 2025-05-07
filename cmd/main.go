package main

import (
	"github.com/Reversive/jindo/internal/handlers"
	"github.com/Reversive/jindo/internal/repositories"
	"github.com/Reversive/jindo/internal/router"
	"github.com/Reversive/jindo/internal/services"
)

func main() {
	r := router.NewRouter()
	novelRepository := repositories.NewNovelRepository()
	novelService := services.NewNovelService(novelRepository)
	novelHandler := handlers.NewNovelHandler(novelService)

	r.Get("/", novelHandler.GetNovels)

	r.Run(":8080")
	return
}
