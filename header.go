package ghc

func (gc *GhcClient) SetHeader(k, v string) *GhcClient {
	gc.Header[k] = v
	return gc
}
