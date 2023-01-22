package port

import (
	"context"
	"jang-article/internal/model"
)

type ArticleUsecase interface {
	SaveArticle(ctx context.Context, article model.Article) (model.Article, error)
}
