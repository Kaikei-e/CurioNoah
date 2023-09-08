package oath

import "golang.org/x/oauth2"

func GetOathConfig() oauth2.Config {
	oathConfig := oauth2.Config{
		ClientID:     "dummy-id",
		ClientSecret: "dummy-secret",
		Scopes:       []string{"dummy1", "dummy2"},
		RedirectURL:  "http://curionoah:4173/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://curionoah.com/auth",
			TokenURL: "http://curionoah.com/token",
		},
	}

	return oathConfig
}
