package repositories

import (
	"context"

	"github.com/Reversive/jindo/internal/models"
)

type NovelRepository interface {
	Create(ctx context.Context, novel *models.Novel) (*models.Novel, error)
	Close()
}
