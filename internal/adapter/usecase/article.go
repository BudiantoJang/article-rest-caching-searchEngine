package usecase

import (
	"context"
	"fmt"
	"jang-article/internal/model"
	"jang-article/internal/port"
)

type article usecase

func (ucs *Usecases) GetArticleUsecase() port.ArticleUsecase {
	return ucs.Article
}

func (uc *article) SaveArticle(ctx context.Context, article model.Article) (model.Article, error) {
	err := uc.ucs.validator.ValidateRequest(ctx, article)
	if err != nil {
		fmt.Println(err)
		return model.Article{}, err
	}
	// can be improved by checking jwt token to verify user (author must be logged in)

	out, err := uc.ucs.database.Save(article)
	if err != nil {
		return model.Article{}, err
	}

	return out, nil
}
