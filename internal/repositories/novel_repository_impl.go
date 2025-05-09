package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/Reversive/jindo/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NovelRepositoryImpl struct {
	Pool *pgxpool.Pool
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

func NewNovelRepository(url string) NovelRepository {
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("Database connection established successfully\n")
	return &NovelRepositoryImpl{Pool: dbpool}
}

func (nri *NovelRepositoryImpl) Close() {
	nri.Pool.Close()
}
