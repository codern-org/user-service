package port

import "github.com/codern-org/user-service/pkg/model"

type UserRepository interface {
	Create(user *model.User) error
	Get(id string) (*model.User, error)
	GetBySessionId(id string) (*model.User, error)
	GetSelfProviderUser(email string) (*model.User, error)
}

type UserService interface {
	HashId(id string, provider model.AuthProvider) string
	Create(email string, password string) (*model.User, error)
	CreateFromGoogle(id string, email string) (*model.User, error)
	Get(id string) (*model.User, error)
	GetBySessionId(id string) (*model.User, error)
	GetSelfProviderUser(email string) (*model.User, error)
}
