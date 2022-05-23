package vismaonline

import (
	"golang.org/x/oauth2"
)

const ()

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://identity.vismaonline.com/connect/authorize",
				TokenURL:  "https://identity.vismaonline.com/connect/token",
				AuthStyle: oauth2.AuthStyleInHeader,
			},
		},
	}

	return config
}
