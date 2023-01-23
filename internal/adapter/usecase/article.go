package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"jang-article/internal/model"
	"jang-article/internal/port"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type article usecase

func (ucs *Usecases) GetArticleUsecase() port.ArticleUsecase {
	return ucs.Article
}

func (uc *article) SaveArticle(ctx context.Context, article model.Article) (model.Article, error) {
	err := uc.ucs.validator.ValidateRequest(ctx, article)
	if err != nil {
		return model.Article{}, err
	}
	// can be improved by checking jwt token to verify user (author must be logged in)

	out, err := uc.ucs.database.Save(article)
	if err != nil {
		return model.Article{}, err
	}

	var docs []redisearch.Document

	docs = append(docs, redisearch.NewDocument(article.Title, 1.0).
		Set("title", article.Title).
		Set("author", article.Author).
		Set("body", article.Body))

	err = uc.ucs.search.UpdateIndex(docs)
	if err != nil {
		return model.Article{}, err
	}

	return out, nil
}

func (uc *article) GetArticle(ctx context.Context, searchparam string) ([]model.Article, error) {
	var out []model.Article

	redisKey := fmt.Sprintf("get_article:" + searchparam)
	// can be improved by checking jwt token to verify user (author must be logged in)

	cache, err := uc.ucs.cache.Get(ctx, redisKey)
	if err != nil {
		if searchparam == "" {
			out, err := uc.ucs.database.GetAll()
			if err != nil {
				return out, err
			}

			_, err = uc.ucs.cache.SetIfNotExist(ctx, redisKey, out, 10*time.Minute)
			if err != nil {
				return out, nil
			}
			return out, nil
		}

		articleTitle, err := uc.ucs.search.GetArticle(ctx, searchparam)
		if err != nil {
			return []model.Article{}, nil
		}

		if articleTitle == "" {
			return []model.Article{}, errors.New("no document matched the searched parameter")
		}

		out, err = uc.ucs.database.FindByTitle(articleTitle)
		if err != nil {
			return []model.Article{}, err
		}
		_, err = uc.ucs.cache.SetIfNotExist(ctx, redisKey, out, 10*time.Minute)
		if err != nil {
			return out, nil
		}
		return out, nil
	}

	err = json.Unmarshal([]byte(cache), &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
