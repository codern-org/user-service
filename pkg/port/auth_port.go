package port

import "github.com/codern-org/user-service/pkg/model"

type AuthService interface {
	Authenticate(header string) (*model.User, error)
	SignIn(email string, password string, ipAddress string, userAgent string) (string, error)
	SignInWithGoogle(code string, ipAddress string, userAgent string) (string, error)
	SignOut(header string) error
}
