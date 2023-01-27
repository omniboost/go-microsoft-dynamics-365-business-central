package central_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	central "github.com/omniboost/go-microsoft-dynamics-365-business-central"
	"golang.org/x/oauth2"
)

var (
	client *central.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")

	var baseURL *url.URL
	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	oauthConfig := central.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	// b, _ := json.MarshalIndent(oauthConfig, "", "  ")
	// log.Fatal(string(b))

	token := &oauth2.Token{
		// AccessToken:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6Ii1LSTNROW5OUjdiUm9meG1lWm9YcWJIWkdldyIsImtpZCI6Ii1LSTNROW5OUjdiUm9meG1lWm9YcWJIWkdldyJ9.eyJhdWQiOiJodHRwczovL2FwaS5idXNpbmVzc2NlbnRyYWwuZHluYW1pY3MuY29tIiwiaXNzIjoiaHR0cHM6Ly9zdHMud2luZG93cy5uZXQvZjc5M2RhNjYtODQwMS00MmVkLThlNWQtMjliOTI5NTkwOTNlLyIsImlhdCI6MTY3MzQ0MTgzNywibmJmIjoxNjczNDQxODM3LCJleHAiOjE2NzM0NDY5OTAsImFjciI6IjEiLCJhaW8iOiJBVFFBeS84VEFBQUFNZ2g0ZXhKSFNhajYvZU1oM0FhWlN4WTdUZWw4dkMyeUZweE9VMUVBbmRmbWFOSmYyMDhCQnRtdDgxalZZcFc1IiwiYW1yIjpbInB3ZCJdLCJhcHBpZCI6ImU0NGI5NjRmLTdhNDItNDE0Mi04OTY0LTdmYzg0Mjg4NjI2OSIsImFwcGlkYWNyIjoiMSIsImZhbWlseV9uYW1lIjoiQm9nYWVydCIsImdpdmVuX25hbWUiOiJMZW9uIiwiaXBhZGRyIjoiNjIuMjM4LjEwNy4yMzMiLCJuYW1lIjoiTGVvbiBCb2dhZXJ0Iiwib2lkIjoiZDBjZjM5MzgtYmQ3NC00YzVjLTliZTQtZTQwZDAwNDcyYTg0IiwicHVpZCI6IjEwMDMzRkZGODI2NTZDQjAiLCJyaCI6IjAuQVFJQVp0cVQ5d0dFN1VLT1hTbTVLVmtKUGozdmJabHNzMU5CaGdlbV9Ud0J1SjhDQUZJLiIsInNjcCI6IkZpbmFuY2lhbHMuUmVhZFdyaXRlLkFsbCB1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJEa3RjcXlMZlc5c3dDN3Y4X1VwTWo1Y3hKekFSRXp6UzV4dVRuTkd0Z2JvIiwidGlkIjoiZjc5M2RhNjYtODQwMS00MmVkLThlNWQtMjliOTI5NTkwOTNlIiwidW5pcXVlX25hbWUiOiJsZW9uQHRpbS1vbmxpbmUubmwiLCJ1cG4iOiJsZW9uQHRpbS1vbmxpbmUubmwiLCJ1dGkiOiI1bGlfN21uWEprV08wTFpJazBodEFBIiwidmVyIjoiMS4wIiwid2lkcyI6WyI2MmU5MDM5NC02OWY1LTQyMzctOTE5MC0wMTIxNzcxNDVlMTAiLCJiNzlmYmY0ZC0zZWY5LTQ2ODktODE0My03NmIxOTRlODU1MDkiXX0.GzxiiOqnlYsBKALr0yrc5Cgj8btcXf92ttbuULQK_6pS2b3GnrxvU6NufkR6peB1n4wriMqR_agKem5E7g4UE3OJ487OH8E_-tBQmuwFAJIaprVlkVlR_5H0NJsGVWHOa-bjlvm8m5o59woHq4Pa26xsh_FiifbGkmbAuOs3NgI0lHPRDPYyRe5tC-CbJczcB6pHwg66yZsTWKlSHH5cRzbQBzpYhEJZ0lL0zJMRaOeyH8v1kt71yerEY2-HEC9dAIpir_EkN0WG_L8ERvO4b6i5vsSHT4k5eHxgDFtLYpJ6jluuUkze9HarwNz0RbC6ZiOKW-mYNbQnpxX2lZsXPQ",
		// Expiry:       time.Now().Add(60 * time.Minute),
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client = central.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
