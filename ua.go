package ghc

import "strings"

func (gc *GhcClient) SetUserAgent(ua string) *GhcClient {
	gc.Header["User-Agent"] = ua
	return gc
}

func (gc *GhcClient) GenerateUserAgent(b string) string {
	ua := "GhcClient/1.0"
	switch strings.ToLower(b) {
	case "chrome":
		ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	case "firefox":
		ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0"
	}
	return ua
}
