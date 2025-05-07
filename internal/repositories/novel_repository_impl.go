package repositories

import (
	"context"

	"github.com/Reversive/jindo/internal/models"
)

type NovelRepositoryImpl struct {
	str string
}

func (nri *NovelRepositoryImpl) Create(
	ctx context.Context,
	novel *models.Novel,
) (*models.Novel, error) {
	return &models.Novel{
		ID:            1,
		Title:         novel.Title,
		Author:        novel.Author,
		Description:   novel.Description,
		Genres:        novel.Genres,
		PublishedDate: novel.PublishedDate,
		CoverImageURL: novel.CoverImageURL,
	}, nil
}

func NewNovelRepository() NovelRepository {
	return &NovelRepositoryImpl{str: "Hello world\n"}
}
