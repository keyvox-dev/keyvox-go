package article

import (
	"fmt"
	"keyvox/util"
)

type Article struct {
	BaseURL string
	APIKey  string
}

func NewArticle(baseURL, apiKey string) *Article {
	return &Article{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (t *Article) List() (string, error) {
	url := fmt.Sprintf("%s/articles", t.BaseURL)
	data, err := util.FetchData(url, t.APIKey)
	if err != nil {
		return "", fmt.Errorf("error fetching articles list: %s", err)
	}

	return fmt.Sprintf("%v", data), nil
}

func (t *Article) Retrieve(idOrSlug string) (string, error) {
	url := fmt.Sprintf("%s/articles/%s", t.BaseURL, idOrSlug)
	data, err := util.FetchData(url, t.APIKey)
	if err != nil {
		return "", fmt.Errorf("error fetching article: %s", err)
	}
	return fmt.Sprintf("%v", data), nil
}
