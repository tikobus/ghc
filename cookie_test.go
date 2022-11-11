package ghc

import "testing"

func TestSetCookie(t *testing.T) {
	client := NewClient()
	client.SetCookie("user", "admin")
	userCookie := ""
	for k, v := range client.Cookie {
		if k == "user" {
			userCookie = v
		}
	}
	if userCookie != "admin" {
		t.Fatal("Set Cookie Failed")
	}
}
