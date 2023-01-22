package usecase

import (
	"jang-article/internal/model"
	"jang-article/internal/port"
)

type article usecase

func (ucs *Usecases) GetArticleUsecase() port.ArticleUsecase {
	return ucs.Article
}

func (uc *article) CreateNewArticle(article model.Article) (model.Article, error) {
	out, err := uc.ucs.database.Save(article)
	if err != nil {
		return model.Article{}, err
	}

	return out, nil
}
