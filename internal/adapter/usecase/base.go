package usecase

import (
	"jang-article/internal/port"
)

type (
	Usecases struct {
		Article *article

		common    usecase
		validator port.Validation
		database  port.DatabaseRepository
		cache     port.CacheRepository
	}

	usecase struct {
		ucs *Usecases
	}
)

func NewUsecases(v port.Validation, pg port.DatabaseRepository, rd port.CacheRepository) *Usecases {
	uc := &Usecases{
		validator: v,
		database:  pg,
		cache:     rd,
	}

	uc.common.ucs = uc
	uc.Article = (*article)(&uc.common)

	return uc
}
