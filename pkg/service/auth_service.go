package service

import (
	"errors"

	"github.com/codern-org/user-service/pkg/model"
	"github.com/codern-org/user-service/pkg/port"
	"golang.org/x/crypto/bcrypt"
)

type AuthSerivce struct {
	googleService  port.GoogleService
	sessionService port.SessionService
	userService    port.UserService
}

func NewAuthService(
	googleService port.GoogleService, sessionService port.SessionService, userService port.UserService,
) *AuthSerivce {
	return &AuthSerivce{
		googleService:  googleService,
		sessionService: sessionService,
		userService:    userService,
	}
}

func (service *AuthSerivce) Authenticate(header string) (*model.User, error) {
	_, error := service.sessionService.Validate(header)
	if error != nil {
		return nil, error
	}
	return nil, nil
}

func (service *AuthSerivce) SignIn(
	email string, password string, ipAddress string, userAgent string,
) (string, error) {
	user, error := service.userService.GetSelfProviderUser(email)
	if error != nil {
		return "", error
	}
	if user == nil {
		return "", errors.New("cannot retrieve self provider user data")
	}

	error = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if error != nil {
		return "", errors.New("password is incorrect")
	}

	return service.sessionService.Create(user.Id, ipAddress, userAgent)
}

func (service *AuthSerivce) SignInWithGoogle(
	code string, ipAddress string, userAgent string,
) (string, error) {
	token, error := service.googleService.GetToken(code)
	if error != nil {
		return "", error
	}
	googleUser, error := service.googleService.GetUser(token)
	if error != nil {
		return "", error
	}

	userId := service.userService.HashId(googleUser.Id, model.GOOGLE)
	user, error := service.userService.Get(userId)
	if error != nil {
		return "", error
	}

	if user == nil {
		user, error = service.userService.CreateFromGoogle(googleUser.Id, googleUser.Email)
		if error != nil {
			return "", error
		}
	}

	return service.sessionService.Create(user.Id, ipAddress, userAgent)
}

func (service *AuthSerivce) SignOut(header string) error {
	session, error := service.sessionService.Validate(header)
	if error != nil {
		return error
	}
	service.sessionService.Destroy(session.Id)
	return nil
}
