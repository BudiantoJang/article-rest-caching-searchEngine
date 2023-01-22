package usecase

import "jang-article/internal/port"

type Usecases struct {
	Article *article

	common   usecase
	database port.DatabaseRepository
	cache    port.CacheRepository
}

type usecase struct {
	ucs *Usecases
}

func NewUsecases(pg port.DatabaseRepository, rd port.CacheRepository) *Usecases {
	uc := &Usecases{
		database: pg,
		cache:    rd,
	}

	uc.common.ucs = uc
	uc.Article = (*article)(&uc.common)

	return uc
}
