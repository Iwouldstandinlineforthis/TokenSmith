package auth

import (
	"fmt"
	"net/url"
    "encoding/json"

	"github.com/gladstomych/tokensmith/internal/classes"
)



func buildAuthURL(clientID, scope, redirectURI string) string {

	// Using net/url to safely build the URL and encode parameters
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientID)
	params.Add("redirect_uri", redirectURI)
	params.Add("scope", scope)

	return fmt.Sprintf("%s?%s", classes.AuthorizeV2Endpoint, params.Encode())
}

func parseRespTokens(respBody string) (*classes.TokenResponse, error) {
    var tr classes.TokenResponse
    err := json.Unmarshal([]byte(respBody), &tr)
    if err != nil {
        return nil, fmt.Errorf("response is not valid JSON: %w", err)
    }

    // Check all required fields
    if tr.RefreshToken == "" || tr.AccessToken == "" || tr.Scope == "" {
        return nil, fmt.Errorf("Token parsing error: missing one or more required fields (refresh_token, access_token, scope). Raw response: %s", respBody)
    }

    fmt.Println("\n[+] SUCCESSFULLY REDEEMED TOKENS!")
    // Print the tokens nicely
    // fmt.Println(tr.AccessToken, "\n")
    fmt.Printf("\n[+] Access Token: \n============================\n%s\n",tr.AccessToken)
    // fmt.Println(tr.RefreshToken, "\n")
    fmt.Printf("\n[+] Refresh Token: \n=============================\n%s\n",tr.RefreshToken)
    // fmt.Println(tr.Scope)
    fmt.Printf("\n[+] Scope: \n=============================\n%s\n", tr.Scope)

    return &tr, nil
}
