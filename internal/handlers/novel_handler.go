package handlers

import (
	"fmt"
	"net/http"

	"github.com/Reversive/jindo/internal/services"
)

type NovelHandler struct {
	novelService *services.NovelService
}

func NewNovelHandler(novelService *services.NovelService) *NovelHandler {
	return &NovelHandler{novelService: novelService}
}

func (nh *NovelHandler) GetNovels(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello from getnovels\n")
}
