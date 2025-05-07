package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Reversive/jindo/internal/dtos"
	"github.com/Reversive/jindo/internal/httpio"
	"github.com/Reversive/jindo/internal/models"
	"github.com/Reversive/jindo/internal/services"
)

type NovelHandler struct {
	novelService *services.NovelService
}

func NewNovelHandler(novelService *services.NovelService) *NovelHandler {
	return &NovelHandler{novelService: novelService}
}

func (nh *NovelHandler) PostNovel(w http.ResponseWriter, req *http.Request) {
	var novelReq dtos.NovelRequest
	err := json.NewDecoder(req.Body).Decode(&novelReq)
	if err != nil {
		// TODO: Discriminate errors from marshalling
		httpio.RespondWithError(w, req, http.StatusBadRequest, err)
		return
	}

	var novel *models.Novel
	novel, err = nh.novelService.CreateNovel(req.Context(), &novelReq)
	if err != nil {
		// TODO: Discriminate errors from the service
		httpio.RespondWithError(w, req, http.StatusConflict, err)
		return
	}

	novelRes := dtos.NovelResponse{
		ID:            novel.ID,
		Title:         novel.Title,
		Author:        novel.Author,
		Description:   novel.Description,
		Genres:        novel.Genres,
		PublishedDate: novel.PublishedDate,
		CoverImageURL: novel.CoverImageURL,
	}

	httpio.RespondWithJSON(w, req, http.StatusCreated, novelRes)
}
