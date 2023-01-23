package model

type SearchArticle struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type SearchArticleOut struct {
	Article []SearchArticle `json:"articles"`
	Total   int             `json:"total"`
}
