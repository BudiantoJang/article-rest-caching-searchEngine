package model

import "time"

type Article struct {
	ID        int
	Author    string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
