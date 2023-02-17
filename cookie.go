package ghc

func (gc *GhcClient) SetCookie(k, v string) *GhcClient {
	gc.Cookie[k] = v
	return gc
}

func (gc *GhcClient) UnsetCookie(k string) *GhcClient {
	delete(gc.Cookie, k)
	return gc
}
