package port

import "github.com/codern-org/user-service/pkg/model"

type GoogleService interface {
	GetOAuthUrl() string
	GetToken(code string) (string, error)
	GetUser(accessToken string) (*model.GoogleUserResponse, error)
}
