package multivers

import (
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

const (
	scope                = "http://UNIT4.Multivers.API/Web/WebApi/*"
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:  BaseURL.String() + "/OAuth/Authorize",
				TokenURL: BaseURL.String() + "/OAuth/Token",
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider(baseURL.String())

	c.Config.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/OAuth/Authorize",
		TokenURL: baseURL.String() + "/OAuth/Token",
	}
}
