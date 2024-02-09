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

func (a *Article) List() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/articles", a.BaseURL)
	data, err := util.FetchData(url, a.APIKey)
	if err != nil {
		return nil, fmt.Errorf("error fetching article list: %s", err)
	}

	articlesData, ok := data["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("data does not contain 'data' field")
	}

	var articles []map[string]interface{}
	for _, article := range articlesData {
		if articleMap, ok := article.(map[string]interface{}); ok {
			articles = append(articles, articleMap)
		} else {
			return nil, fmt.Errorf("invalid article format")
		}
	}

	return articles, nil
}

func (a *Article) Retrieve(idOrSlug string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/articles/%s", a.BaseURL, idOrSlug)
	data, err := util.FetchData(url, a.APIKey)
	if err != nil {
		return nil, fmt.Errorf("error fetching article: %s", err)
	}

	articleData, ok := data["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("article data not found in response")
	}

	return articleData, nil
}
