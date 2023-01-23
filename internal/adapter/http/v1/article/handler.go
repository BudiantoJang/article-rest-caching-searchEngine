package article

import (
	"jang-article/internal/port"

	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	article port.ArticleUsecase
}

// New article handler will initiate the article / resources endpoint
func New(app *echo.Group, art port.ArticleUsecase) {
	ah := articleHandler{
		article: art,
	}

	article := app.Group("/article")
	article.POST("", ah.CreateNewArticle())
	article.GET("", ah.FindArticle())
}
