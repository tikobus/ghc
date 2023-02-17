package ghc

func (gc *GhcClient) SetHeader(k, v string) *GhcClient {
	gc.Header[k] = v
	return gc
}

func (gc *GhcClient) UnsetHeader(k string) *GhcClient {
	delete(gc.Header, k)
	return gc
}
