package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/codern-org/user-service/pkg/model"
)

type GoogleService struct {
	httpClient   *http.Client
	clientId     string
	clientSecert string
	redirectUri  string
}

func NewGoogleService(clientId string, clientSecret string, redirectUri string) *GoogleService {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	return &GoogleService{
		httpClient:   httpClient,
		clientId:     clientId,
		clientSecert: clientSecret,
		redirectUri:  redirectUri,
	}
}

func (service *GoogleService) GetOAuthUrl() string {
	query := url.Values{}
	query.Add("redirect_uri", service.redirectUri)
	query.Add("client_id", service.clientId)
	query.Add("access_type", "offline")
	query.Add("response_type", "code")
	query.Add("prompt", "consent")
	query.Add("scope", strings.Join([]string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}, " "))
	return "https://accounts.google.com/o/oauth2/v2/auth?" + query.Encode()
}

func (service *GoogleService) GetToken(code string) (string, error) {
	body, error := json.Marshal(&map[string]string{
		"code":          code,
		"client_id":     service.clientId,
		"client_secret": service.clientSecert,
		"redirect_uri":  service.redirectUri,
		"grant_type":    "authorization_code",
	})
	if error != nil {
		return "", error
	}

	request, error := http.NewRequest("POST", "https://oauth2.googleapis.com/token", bytes.NewReader(body))
	if error != nil {
		return "", error
	}

	response, error := service.httpClient.Do(request)
	if error != nil {
		return "", error
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("cannot get token from google api, status code: " + response.Status)
	}

	data, error := io.ReadAll(response.Body)
	if error != nil {
		return "", error
	}

	var result model.GoogleTokenResponse
	if error = json.Unmarshal(data, &result); error != nil {
		return "", error
	}

	return result.AccessToken, nil
}

func (service *GoogleService) GetUser(accessToken string) (*model.GoogleUserResponse, error) {
	request, error := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if error != nil {
		return nil, error
	}

	response, error := service.httpClient.Do(request)
	if error != nil {
		return nil, error
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("cannot get user from google api, status code: " + response.Status)
	}

	data, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	var result model.GoogleUserResponse
	if error = json.Unmarshal(data, &result); error != nil {
		return nil, error
	}

	return &result, error
}
