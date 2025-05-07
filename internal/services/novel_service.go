package services

import "github.com/Reversive/jindo/internal/repositories"

type NovelService struct {
	novelRepository repositories.NovelRepository
}

func NewNovelService(
	novelRepository repositories.NovelRepository,
) *NovelService {
	return &NovelService{novelRepository: novelRepository}
}
