package pg

import (
	"jang-article/internal/model"
	"jang-article/internal/port"

	"gorm.io/gorm"
)

type Client struct {
	Client *gorm.DB
}

func New(client *gorm.DB) port.DatabaseRepository {
	return &Client{client}
}

func (p *Client) Save(article model.Article) (model.Article, error) {
	err := p.Client.Save(&article).Error
	if err != nil {
		return model.Article{}, err
	}

	return article, nil
}

func (p Client) GetAll() ([]model.Article, error) {
	var articles []model.Article

	err := p.Client.Find(&articles).Limit(5).Error
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}

func (p *Client) FindByAuthor(author string) ([]model.Article, error) {
	var article []model.Article

	err := p.Client.Where("author=?", author).Find(&article).Limit(5).Error
	if err != nil {
		return []model.Article{}, err
	}

	return article, nil
}

func (p *Client) FindByTitle(title string) ([]model.Article, error) {
	var article []model.Article

	err := p.Client.Where("title=?", title).Find(&article).Limit(5).Error
	if err != nil {
		return []model.Article{}, err
	}

	return article, nil
}
