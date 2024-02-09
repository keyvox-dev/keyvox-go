package keyvox

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var kv *KeyVox

func Setup(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf("error loading .env file: %s", err)
	}

	apiKey := os.Getenv("API_KEY")
	baseURL := os.Getenv("BASE_URL")

	kv = NewKeyVox(apiKey, baseURL)
}

func TestNewKeyVox(t *testing.T) {
	Setup(t)
}

func TestArticlesList(t *testing.T) {
	Setup(t)
	articles, err := kv.articles.List()
	if err != nil {
		t.Fatalf("error listing articles: %s", err)
	}
	for _, article := range articles {
		fmt.Println(article)
	}
	fmt.Println("---------------")
}

func TestArticlesRetrieve(t *testing.T) {
	Setup(t)

	id := os.Getenv("ARTICLE_ID")
	article, err := kv.articles.Retrieve(id)
	if err != nil {
		t.Fatalf("error getting article: %s", err)
	}
	fmt.Println(article)
}
