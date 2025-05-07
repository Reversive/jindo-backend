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

	r.Post("/api/novels", novelHandler.PostNovel)

	r.Run(":8080")
	return
}
