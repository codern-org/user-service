package repository

import (
	"github.com/codern-org/user-service/pkg/model"
	"github.com/jmoiron/sqlx"
)

type SessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (repository *SessionRepository) Create(session *model.Session) error {
	_, error := repository.db.NamedExec(
		"INSERT INTO session VALUES (:id, :user_id, :ip_address, :user_agent, :expired_at, :created_at)",
		session,
	)
	if error != nil {
		return error
	}
	return nil
}

func (repository *SessionRepository) Get(id string) (*model.Session, error) {
	session := model.Session{}
	error := repository.db.Get(&session, "SELECT * FROM session WHERE id = ?", id)
	if error != nil {
		return nil, error
	}
	return &session, nil
}

func (repository *SessionRepository) Delete(id string) error {
	_, error := repository.db.Exec("DELETE FROM session WHERE id = ?", id)
	if error != nil {
		return error
	}
	return nil
}

func (repository *SessionRepository) DeleteDuplicates(userId string, ipAddress string, userAgent string) error {
	_, error := repository.db.Exec(
		"DELETE FROM session WHERE user_id = ? AND user_agent = ? AND ip_address = ?",
		userId, userAgent, ipAddress,
	)
	if error != nil {
		return error
	}
	return nil
}
