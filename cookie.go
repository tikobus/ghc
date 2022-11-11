package ghc

func (gc *GhcClient) SetCookie(k, v string) *GhcClient {
	gc.Cookie[k] = v
	return gc
}
