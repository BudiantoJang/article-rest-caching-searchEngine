package model

import "time"

type Article struct {
	ID        int       `json:"id"`
	Author    string    `json:"author" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateArticleIn struct {
	Author string `json:"author" validate:"required,email"`
	Title  string `json:"title" validate:"required,email"`
	Body   string `json:"body" validate:"required"`
}
