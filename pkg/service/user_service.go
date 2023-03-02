package service

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/mail"

	"github.com/codern-org/user-service/pkg/model"
	"github.com/codern-org/user-service/pkg/port"
	"github.com/codern-org/user-service/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(userRepository port.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) HashId(id string, provider model.AuthProvider) string {
	sha1 := sha1.New()
	sha1.Write([]byte(id + "." + string(provider)))
	return hex.EncodeToString(sha1.Sum(nil))
}

func (service *UserService) Create(email string, password string) (*model.User, error) {
	if _, error := mail.ParseAddress(email); error != nil {
		return nil, fmt.Errorf("email %s is invalid", email)
	}

	user, error := service.GetSelfProviderUser(email)
	if error != nil {
		return nil, error
	}
	if user != nil {
		return nil, fmt.Errorf("email %s already registered", email)
	}

	id := service.HashId(email, model.SELF)
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), 10)
	if error != nil {
		return nil, error
	}

	// TODO: profile generation

	user = &model.User{
		Id:          id,
		Email:       email,
		Password:    string(hashedPassword),
		DisplayName: email,
		ProfilePath: "",
		Provider:    model.SELF,
		CreatedAt:   util.GetTimestamp(),
	}

	if error = service.userRepository.Create(user); error != nil {
		return nil, error
	}
	return user, nil
}

func (service *UserService) CreateFromGoogle(id string, email string) (*model.User, error) {
	// TODO: profile generation

	user := &model.User{
		Id:          service.HashId(id, model.GOOGLE),
		Email:       email,
		Password:    "",
		DisplayName: email,
		ProfilePath: "",
		Provider:    model.SELF,
		CreatedAt:   util.GetTimestamp(),
	}

	if error := service.userRepository.Create(user); error != nil {
		return nil, error
	}
	return user, nil
}

func (service *UserService) Get(id string) (*model.User, error) {
	user, error := service.userRepository.Get(id)
	if error != nil {
		return nil, error
	}
	return user, nil
}

func (service *UserService) GetBySessionId(id string) (*model.User, error) {
	user, error := service.userRepository.GetBySessionId(id)
	if error != nil {
		return nil, error
	}
	return user, nil
}

func (service *UserService) GetSelfProviderUser(email string) (*model.User, error) {
	user, error := service.userRepository.GetSelfProviderUser(email)
	if error != nil {
		return nil, error
	}
	return user, nil
}
