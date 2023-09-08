package oath

import (
	"golang.org/x/oauth2"
	"net/http"
)

func handleAuth(writer http.ResponseWriter, request *http.Request) {
	config := GetOathConfig()
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(writer, request, authURL, http.StatusFound)
}
