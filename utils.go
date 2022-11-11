package ghc

import "time"

func (gc *GhcClient) TimeNow() string {
	return time.Now().String()
}
