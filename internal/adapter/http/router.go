package http

import (
	"context"
	"fmt"
	v1Article "jang-article/internal/adapter/http/v1/article"
	"jang-article/internal/model"
	"jang-article/internal/port"
	"log"

	"github.com/labstack/echo/v4"
)

type svc struct {
	e    *echo.Echo
	addr string
}

func (s *svc) Start(ctx context.Context) {
	if err := s.e.Start(s.addr); err != nil {
		log.Panic(err)
	}
}

func NewRouter(conf model.Config, redis port.CacheRepository, pg port.DatabaseRepository, uc port.Usecases) *svc {
	app := echo.New()
	svc := &svc{
		e:    app,
		addr: fmt.Sprintf(":%d", conf.App.Port),
	}

	v1Group := app.Group("/api/v1")

	v1Article.New(v1Group, uc.GetArticleUsecase())

	return svc
}
