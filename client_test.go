package ghc

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client.Timeout == 0 {
		t.Fatal("error")
	}
}
