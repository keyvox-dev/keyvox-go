package keyvox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type KeyVox struct {
	APIKey  string
	BaseURL string
}

func NewKeyVox(apiKey string) *KeyVox {
	return &KeyVox{
		APIKey:  apiKey,
		BaseURL: "https://keyvox.dev/api",
	}
}

func (kv *KeyVox) FetchData(url string) (map[string]interface{}, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %s", err)
	}

	// Add the API key to the request headers
	req.Header.Set("key", kv.APIKey)

	// Perform the HTTP request
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing HTTP request: %s", err)
	}
	defer response.Body.Close()

	// Read the body of the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	// Decode the JSON data into a map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON data: %s", err)
	}

	return data, nil
}
