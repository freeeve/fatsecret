package fatsecret

import (
	"github.com/mrjones/oauth"
)

type FatSecretConn struct {
	Consumer *oauth.Consumer
}

func Connect(apikey, secret string) (FatSecretConn, error) {
	return FatSecretConn{Consumer: oauth.NewConsumer(
		apikey,
		secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://platform.fatsecret.com/rest/server.api",
			AuthorizeTokenUrl: "http://platform.fatsecret.com/rest/server.api",
			AccessTokenUrl:    "http://platform.fatsecret.com/rest/server.api",
		})}, nil
}
