package port

import "jang-article/internal/model"

type DatabaseRepository interface {
	Save(article model.Article) (model.Article, error)
	GetAll() ([]model.Article, error)
	FindByAuthor(author string) ([]model.Article, error)
	FindByTitle(title string) ([]model.Article, error)
}
