package main

import (
	"fmt"
	"ghc"
)

func main() {
	client := ghc.NewClient()
	client.SetUrl("https://github.com/manifest.json")
	client.SetMethod("GET")
	client.SetHeader("Host", "github.com")
	client.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	client.SetCookie("user", "123456")
	client.SetKeepalive(true)
	client.SetTimeout(20)
	client.DoAll()
	fmt.Print(client.GetResponseBodyString())
}
