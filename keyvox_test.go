package keyvox

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var kv *KeyVox

func Setup(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf(".env not found: %s", err)
	}

	apiKey := os.Getenv("API_KEY")

	kv := NewKeyVox(apiKey)
	kv.BaseURL = os.Getenv("BASE_URL")
}

func TestNewKeyVox(t *testing.T) {
	Setup(t)
}

func TestArticlesList(t *testing.T) {
	Setup(t)

}
