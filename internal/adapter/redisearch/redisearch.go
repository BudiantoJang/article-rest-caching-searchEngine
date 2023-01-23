package redisearch

import (
	"context"
	"fmt"
	"jang-article/internal/port"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type Client struct {
	Client *redisearch.Client
}

func New(client *redisearch.Client) port.SearchRepository {
	return &Client{client}
}

func (r *Client) Info() (*redisearch.IndexInfo, error) {
	return r.Client.Info()
}

func (r *Client) CreateSchema(schema *redisearch.Schema) error {
	return r.Client.CreateIndex(schema)
}

func (r *Client) UpdateIndex(docs []redisearch.Document) error {
	indexingOptions := redisearch.IndexingOptions{
		Replace: true,
	}

	return r.Client.IndexOptions(indexingOptions, docs...)
}

func (r *Client) GetArticle(ctx context.Context, in string) (string, error) {
	searchParam := fmt.Sprintf("*%s*", in)
	docs, _, err := r.Client.Search(redisearch.NewQuery(searchParam).
		Limit(0, 1).
		SetReturnFields("author", "title", "body"))

	if err != nil || len(docs) == 0 {
		return "", err
	}

	searchResult := fmt.Sprintf("%v", docs[0].Properties["title"])
	return searchResult, nil

}
