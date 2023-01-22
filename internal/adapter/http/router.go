package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type svc struct {
	e *echo.Echo
	addr string
}

func NewRouter(conf model.Cofig, redis port.CacheRepository, redisearch port.SearchRepository) *svc{
	app := echo.New()
	svc := &svc{
		e: app,
		addr: fmt.Sprintf("%d", conf.App.Port)
	}
}