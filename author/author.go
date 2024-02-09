package author

import (
	"fmt"
	"keyvox/util"
)

type Author struct {
	BaseURL string
	APIKey  string
}

func NewAuthor(baseURL, apiKey string) *Author {
	return &Author{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (t *Author) List() (string, error) {
	url := fmt.Sprintf("%s/authors", t.BaseURL)
	data, err := util.FetchData(url, t.APIKey)
	if err != nil {
		return "", fmt.Errorf("error fetching authors list: %s", err)
	}

	return fmt.Sprintf("%v", data), nil
}

func (t *Author) Retrieve(idOrSlug string) (string, error) {
	url := fmt.Sprintf("%s/authors/%s", t.BaseURL, idOrSlug)
	data, err := util.FetchData(url, t.APIKey)
	if err != nil {
		return "", fmt.Errorf("error fetching author: %s", err)
	}
	return fmt.Sprintf("%v", data), nil
}
