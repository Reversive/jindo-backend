package dtos

import (
	"time"
)

type NovelRequest struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverImageURL string    `json:"cover_image_url"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
	Genres        []string  `json:"genres,omitempty"`
}

type NovelResponse struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverImageURL string    `json:"cover_image_url"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
	Genres        []string  `json:"genres,omitempty"`
}
