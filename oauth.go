package poweroffice

import (
	"net/url"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const ()

type Oauth2Config struct {
	clientcredentials.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{},
			TokenURL:     BaseURL.String() + "/OAuth/Token",
			AuthStyle:    oauth2.AuthStyleInHeader,
		},
	}

	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider(baseURL.String())

	c.Config.TokenURL = baseURL.String() + "/OAuth/Token"
}
