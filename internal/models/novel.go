package models

import "time"

type Novel struct {
	ID            int
	Title         string
	Author        string
	Description   string
	CoverImageURL string
	Genres        []string
	PublishedDate time.Time
}
