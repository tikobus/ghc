package ghc

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	client.SetUrl("https://github.com/manifest.json").SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	_, err := client.DoAll()
	if err != nil {
		t.Fatal("Get Response Failed")
	}
	if client.GetResponseStatus() != 200 {
		t.Fatal("Get ResponseStatus Not 200")
	}
}
