package central

import (
	"time"

	"golang.org/x/oauth2"
)

const (
	// scope                = "*"
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
	APIKey string
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			// Scopes:       []string{"https://graph.microsoft.com/.default", "openid", "offline_access", "https://api.businesscentral.dynamics.com/.default"},
			Scopes: []string{"openid", "offline_access", "profile"},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://login.microsoftonline.com/common/oauth2/v2.0/token",
			},
		},
	}

	return config
}
