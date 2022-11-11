package main

import (
	"fmt"
	"ghc"
)

func main() {
	client := ghc.NewClient()
	client.SetUrl("https://www.baidu.com/")
	client.SetMethod("GET")
	client.SetHeader("Host", "www.baidu.com")
	client.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	client.SetCookie("user", "123456")
	client.SetKeepalive(true)
	client.SetTimeout(20)
	resp, _ := client.DoAll()
	fmt.Print(string(resp))
}
