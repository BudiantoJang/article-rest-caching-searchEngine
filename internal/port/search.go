package port

import (
	"context"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type SearchRepository interface {
	Info() (*redisearch.IndexInfo, error)
	CreateSchema(schema *redisearch.Schema) error
	UpdateIndex(docs []redisearch.Document) error
	GetArticle(ctx context.Context, in string) (string, error)
}
