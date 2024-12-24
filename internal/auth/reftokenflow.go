package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gladstomych/tokensmith/internal/classes"
)

func GetAccTknFromRefTkn() {
	fmt.Println("Obtaining New Access Tokens from Refresh tokens.")

	clientID := classes.ClientID
	resourceURL := classes.RefResourceURL
	scope := classes.Scope
	refToken := classes.RefreshToken
	userAgent := classes.UserAgent

	// redeemedHTTPResp := reqAccessTknFlow(clientID, scope, refToken, resourceURL, userAgent)
	err := reqAccessTknFlow(clientID, scope, refToken, resourceURL, userAgent)
	if err != nil {
		log.Fatalln(err)
	}
}

func reqAccessTknFlow(clientID, scope, refToken, resource, userAgent string) error {

	tokenURL := classes.TokenV1Endpoint
	body := fmt.Sprintf("client_id=%s&grant_type=refresh_token&scope=%s&resource=%s&refresh_token=%s", clientID, scope, resource, refToken)

	Client := &http.Client{}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//var ptr *classes.TokenResponse
	err = parseRespTokens(string(respBody))
	if err != nil {
		return err
	}

	return nil

}

// GOAL:
// POST /common/oauth2/token HTTP/1.1
// Host: login.microsoftonline.com
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36
// Content-Type: application/x-www-form-urlencoded
// Content-Length: 1067

// grant_type=refresh_token&client_id=04b07795-8ddb-461a-bbee-02f9e1bf7b46&scope=openid%20offline_access&resource=https://graph.microsoft.com/&refresh_token=0.AR
