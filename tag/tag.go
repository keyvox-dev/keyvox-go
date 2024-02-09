package tag

import (
	"fmt"
	"keyvox/util"
)

type Tag struct {
	BaseURL string
	APIKey  string
}

func NewTag(baseURL, apiKey string) *Tag {
	return &Tag{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (t *Tag) List() (string, error) {
	url := fmt.Sprintf("%s/tags", t.BaseURL)
	data, err := util.FetchData(url, t.APIKey)
	if err != nil {
		return "", fmt.Errorf("error fetching tag list: %s", err)
	}

	return fmt.Sprintf("%v", data), nil
}

func (t *Tag) Retrieve(idOrSlug string) (string, error) {
	url := fmt.Sprintf("%s/tags/%s", t.BaseURL, idOrSlug)
	data, err := util.FetchData(url, t.APIKey)
	if err != nil {
		return "", fmt.Errorf("error fetching tag: %s", err)
	}
	return fmt.Sprintf("%v", data), nil
}
