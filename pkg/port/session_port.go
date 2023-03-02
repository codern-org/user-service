package port

import "github.com/codern-org/user-service/pkg/model"

type SessionRepository interface {
	Create(session *model.Session) error
	Get(id string) (*model.Session, error)
	Delete(id string) error
	DeleteDuplicates(userId string, ipAddress string, userAgent string) error
}

type SessionService interface {
	Sign(id string) string
	Unsign(header string) (string, error)
	Create(userId string, ipAddress string, userAgent string) (string, error)
	Get(header string) (*model.Session, error)
	Destroy(id string)
	Validate(header string) (*model.Session, error)
}
