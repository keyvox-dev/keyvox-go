package keyvox

import (
	"keyvox/article"
	"keyvox/tag"
)

type KeyVox struct {
	APIKey   string
	BaseURL  string
	articles *article.Article
	tags     *tag.Tag
}

func NewKeyVox(apiKey, baseURL string) *KeyVox {
	kv := &KeyVox{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
	kv.articles = article.NewArticle(baseURL, apiKey)
	kv.tags = tag.NewTag(baseURL, apiKey)
	return kv
}

func (kv *KeyVox) Articles() *article.Article {
	return kv.articles
}

func (kv *KeyVox) Tags() *tag.Tag {
	return kv.tags
}
