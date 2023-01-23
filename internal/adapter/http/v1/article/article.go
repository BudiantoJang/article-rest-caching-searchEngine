package article

import (
	"jang-article/internal/adapter/http/helper"
	"jang-article/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ah *articleHandler) CreateNewArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := model.CreateArticleIn{}

		if err := c.Bind(&req); err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, model.ErrorDetail{
				Code:        model.ResponseCode99,
				Description: err,
			})
		}

		out, err := ah.article.SaveArticle(ctx, model.Article{
			Author: req.Author,
			Title:  req.Title,
			Body:   req.Body,
		})

		if err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, err)
		}

		return helper.ResponseSuccess(c, out)
	}
}

type FindArticleIn struct {
	Query string `query:"query"`
}

func (ah *articleHandler) FindArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		req := FindArticleIn{}
		if err := c.Bind(&req); err != nil {
			helper.ResponseError(c, http.StatusBadRequest, err)
		}

		out, err := ah.article.GetArticle(ctx, req.Query)
		if err != nil {
			return helper.ResponseError(c, http.StatusBadRequest, err)
		}
		return helper.ResponseSuccess(c, out)
	}
}
