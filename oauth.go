package multivers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

type LinkHandler func(*url.URL) error
type AuthorizationHandler func(string) error
type TokenHandler func(*oauth2.Token) error

const (
	localPort                 = "9090"
	authorizationCallbackPath = "/authorization/callback"
	scope                     = "http://UNIT4.Multivers.API/Web/WebApi/*"
	oauthStateString          = ""
	authorizationTimeout      = 60 * time.Second
	tokenTimeout              = 5 * time.Second
)

func NewOauth2Config(baseURL *url.URL) *oauth2.Config {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider(baseURL.String())

	return &oauth2.Config{
		RedirectURL:  getAuthorizationCallbackURL(),
		ClientID:     "",
		ClientSecret: "",
		Scopes:       []string{scope},
		Endpoint: oauth2.Endpoint{
			AuthURL:  baseURL.String() + "/OAuth/Authorize",
			TokenURL: baseURL.String() + "/OAuth/Token",
		},
	}
}

func getAuthorizationCallbackURL() string {
	return fmt.Sprintf("http://localhost:%s%s", localPort, authorizationCallbackPath)
}

func GetNewOauth2Token(oauthConfig *oauth2.Config, linkCallback LinkHandler, authorizationCallback AuthorizationHandler, tokenCallback TokenHandler) (*oauth2.Token, error) {
	url, err := generateLoginLink(oauthConfig)
	if err != nil {
		return nil, err
	}

	// if linkHandler is provided execute it
	if linkCallback != nil {
		err = linkCallback(url)
		if err != nil {
			return nil, err
		}
	}

	authorizationCode, err := waitForAuthorizationCallback()
	if err != nil {
		return nil, err
	}

	// if authorizationHandler is provided execute it
	if authorizationCallback != nil {
		authorizationCallback(authorizationCode)
	}

	token, err := getOauth2Token(oauthConfig, authorizationCode)
	if err != nil {
		return nil, err
	}

	if tokenCallback != nil {
		err := tokenCallback(token)
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

func generateLoginLink(oauthConfig *oauth2.Config) (*url.URL, error) {
	u := oauthConfig.AuthCodeURL(oauthStateString)
	return url.Parse(u)
}

// @TODO: add context so the server can be canceled?
func waitForAuthorizationCallback() (string, error) {
	mux := http.NewServeMux()
	ac := newAuthorizationController()
	mux.HandleFunc("/authorization/callback", ac.callbackHandler)
	srv := &http.Server{Addr: ":" + localPort, Handler: mux}

	go func() {
		// service connections
		srv.ListenAndServe()
	}()

	// block until data is received on one of the channels
	var code string
	var err error
	select {
	case <-time.After(authorizationTimeout):
		err = errors.New("Timed out waiting for authorization callback")
	case code = <-ac.resultCh:
	case err = <-ac.errCh:
	}

	close(ac.errCh)
	close(ac.resultCh)

	return code, err
}

func newAuthorizationController() *authorizationController {
	ac := &authorizationController{}
	ac.errCh = make(chan error)
	ac.resultCh = make(chan string)
	return ac
}

type authorizationController struct {
	errCh    chan error
	resultCh chan string
}

func (c *authorizationController) callbackHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	code := r.URL.Query().Get("code")
	if code == "" {
		log.Println("sending on error channel")
		c.errCh <- fmt.Errorf("No code found in query parameters")
		return
	}
	log.Println("sending on result channel")
	c.resultCh <- code
}

func getOauth2Token(oauthConfig *oauth2.Config, code string) (*oauth2.Token, error) {
	ctx, _ := context.WithTimeout(context.Background(), tokenTimeout)
	return oauthConfig.Exchange(ctx, code)
}
