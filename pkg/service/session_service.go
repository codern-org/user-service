package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/codern-org/user-service/pkg/model"
	"github.com/codern-org/user-service/pkg/port"
	"github.com/codern-org/user-service/pkg/util"
	"github.com/google/uuid"
)

type SessionService struct {
	secretKey         string
	sessionRepository port.SessionRepository
}

const (
	SessionPrefix string = "$"
)

func NewSessionService(
	secretKey string, sessionRepository port.SessionRepository,
) *SessionService {
	return &SessionService{sessionRepository: sessionRepository}
}

func (service *SessionService) Sign(id string) string {
	hmac := hmac.New(sha256.New, []byte(service.secretKey))
	hmac.Write([]byte(id))
	regex := regexp.MustCompile(`=+$`)
	signature := regex.ReplaceAllString(base64.StdEncoding.EncodeToString(hmac.Sum(nil)), "")
	return SessionPrefix + ":" + id + "." + signature
}

func (service *SessionService) Unsign(header string) (string, error) {
	if !strings.HasPrefix(header, SessionPrefix+":") {
		return "", errors.New("prefix mismatch")
	}

	id := header[len(SessionPrefix)+1 : strings.LastIndex(header, ".")]
	expectation := service.Sign(id)

	isLengthMatch := len([]byte(header)) == len([]byte(expectation))
	isInputMatch := subtle.ConstantTimeCompare([]byte(header), []byte(expectation)) == 1

	if !isLengthMatch || !isInputMatch {
		return "", errors.New("signature mismatch")
	}
	return id, nil
}

func (service *SessionService) Create(userId string, ipAddress string, userAgent string) (string, error) {
	error := service.sessionRepository.DeleteDuplicates(userId, userAgent, ipAddress)
	if error != nil {
		return "", error
	}

	id := service.Sign(uuid.NewString())
	maxAge := 7 * 24 * time.Hour

	createdAt := util.GetTimestamp()
	expiredAt := createdAt + int64(maxAge.Seconds())

	error = service.sessionRepository.Create(&model.Session{
		Id:        id,
		UserId:    userId,
		IpAddress: ipAddress,
		UserAgent: userAgent,
		ExpiredAt: expiredAt,
		CreatedAt: createdAt,
	})
	if error != nil {
		return "", error
	}

	return id, nil
}

func (service *SessionService) Get(header string) (*model.Session, error) {
	id, error := service.Unsign(header)
	if error != nil {
		return nil, error
	}

	session, error := service.sessionRepository.Get(id)
	if error != nil {
		return nil, error
	}

	return session, nil
}

func (service *SessionService) Destroy(id string) {
	service.sessionRepository.Delete(id)
}

func (service *SessionService) Validate(header string) (*model.Session, error) {
	session, error := service.Get(header)
	if error != nil {
		return nil, error
	}

	if util.GetTimestamp() >= session.ExpiredAt {
		service.Destroy(session.Id)
		return nil, errors.New("session expired")
	}

	return session, nil
}
