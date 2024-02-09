package keyvox

import "keyvox/article"

type KeyVox struct {
	APIKey   string
	BaseURL  string
	articles *article.Article
}

func NewKeyVox(apiKey, baseURL string) *KeyVox {
	kv := &KeyVox{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
	kv.articles = article.NewArticle(baseURL, apiKey)
	return kv
}

func (kv *KeyVox) Articles() *article.Article {
	return kv.articles
}
