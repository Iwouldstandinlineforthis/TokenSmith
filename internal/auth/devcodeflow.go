package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gladstomych/tokensmith/internal/classes"
)

func reqDeviceCodeFlow(ClientID, resourceURL, userAgent string) (string, error) {
	deviceCodeURL := classes.DeviceCodeEndpoint
	body := fmt.Sprintf("client_id=%s&resource=%s", ClientID, resourceURL)

	Client := &http.Client{}
	req, err := http.NewRequest("POST", deviceCodeURL, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	resp, err := Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	verificationURI, deviceCode, userCode, err := parseDeviceAuthorizationResponse(string(respBody))
	if err != nil {
		return "", err
	}

	fmt.Printf("To authenticate, visit %s and enter the code: %s\n", verificationURI, userCode)
	return deviceCode, nil
}

func pollDeviceCodeFlow(ClientID, deviceCode, userAgent string) error {
	tokenURL := classes.TokenV2Endpoint
	body := fmt.Sprintf("client_id=%s&grant_type=urn:ietf:params:oauth:grant-type:device_code&device_code=%s", ClientID, deviceCode)

	Client := &http.Client{}
	for {
		req, err := http.NewRequest("POST", tokenURL, strings.NewReader(body))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", userAgent)

		resp, err := Client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if strings.Contains(string(respBody), "authorization_pending") {
			fmt.Println("Authorization pending. Retrying...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = parseDeviceAuthRespTokens(string(respBody))
		if err != nil {
			return err
		}

		break
	}

	return nil
}

func GetTknFromDevCode() {
	clientID := classes.ClientID
	resourceURL := classes.ResourceURL

	deviceCode, err := reqDeviceCodeFlow(clientID, resourceURL, classes.UserAgent)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Waiting for user to complete authentication...")

	err = pollDeviceCodeFlow(clientID, deviceCode, classes.UserAgent)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Authentication successful.")
}
