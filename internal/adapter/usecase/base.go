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
		search    port.SearchRepository
		cache     port.CacheRepository
	}

	usecase struct {
		ucs *Usecases
	}
)

func NewUsecases(v port.Validation, pg port.DatabaseRepository, rd port.CacheRepository, rdsearch port.SearchRepository) *Usecases {
	uc := &Usecases{
		validator: v,
		database:  pg,
		cache:     rd,
		search:    rdsearch,
	}

	uc.common.ucs = uc
	uc.Article = (*article)(&uc.common)

	return uc
}
