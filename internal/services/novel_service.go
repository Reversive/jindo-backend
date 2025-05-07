package services

import (
	"context"

	"github.com/Reversive/jindo/internal/dtos"
	"github.com/Reversive/jindo/internal/models"
	"github.com/Reversive/jindo/internal/repositories"
)

type NovelService struct {
	NovelRepository repositories.NovelRepository
}

func NewNovelService(
	novelRepository repositories.NovelRepository,
) *NovelService {
	return &NovelService{NovelRepository: novelRepository}
}

func (ns *NovelService) CreateNovel(
	ctx context.Context,
	novelReq *dtos.NovelRequest,
) (*models.Novel, error) {
	// TODO: Check empty/blank fields (validation/required fields)
	novel := models.Novel{
		Title:         novelReq.Title,
		Author:        novelReq.Author,
		Description:   novelReq.Description,
		Genres:        novelReq.Genres,
		PublishedDate: novelReq.PublishedDate,
		CoverImageURL: novelReq.CoverImageURL,
	}
	return ns.NovelRepository.Create(ctx, &novel)
}
