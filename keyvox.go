package keyvox

import (
	"keyvox/article"
	"keyvox/author"
	"keyvox/tag"
)

type KeyVox struct {
	APIKey   string
	BaseURL  string
	articles *article.Article
	tags     *tag.Tag
	authors  *author.Author
}

func NewKeyVox(apiKey, baseURL string) *KeyVox {
	kv := &KeyVox{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}

	kv.articles = article.NewArticle(baseURL, apiKey)
	kv.tags = tag.NewTag(baseURL, apiKey)
	kv.authors = author.NewAuthor(baseURL, apiKey)

	return kv
}

func (kv *KeyVox) Articles() *article.Article {
	return kv.articles
}

func (kv *KeyVox) Tags() *tag.Tag {
	return kv.tags
}

func (kv *KeyVox) Authors() *author.Author {
	return kv.authors
}
