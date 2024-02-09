package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchData(url string, apiKey string) (map[string]interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %s", err)
	}

	req.Header.Set("key", apiKey)
	req.Header.Set("lang", "go")

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing HTTP request: %s", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %s", err)
	}

	return data, nil

}
