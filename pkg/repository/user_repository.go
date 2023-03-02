package repository

import (
	"github.com/codern-org/user-service/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) Create(user *model.User) error {
	_, error := repository.db.NamedExec(
		"INSERT INTO user (id, email, password, display_name, profile_path, provider, created_at)"+
			"VALUES (:id, :email, :password, :display_name, :profile_path, :provider, :created_at)",
		user,
	)
	if error != nil {
		return error
	}
	return nil
}

func (repository *UserRepository) Get(id string) (*model.User, error) {
	user := model.User{}
	error := repository.db.Get(&user, "SELECT * FROM user WHERE id = ?", id)
	if error != nil {
		return nil, error
	}
	return &user, nil
}

func (repository *UserRepository) GetBySessionId(id string) (*model.User, error) {
	user := model.User{}
	error := repository.db.Get(
		&user,
		"SELECT user.* FROM user JOIN session ON user.id = session.user_id WHERE session.id = ?",
		id,
	)
	if error != nil {
		return nil, error
	}
	return &user, nil
}

func (repository *UserRepository) GetSelfProviderUser(email string) (*model.User, error) {
	user := model.User{}
	error := repository.db.Get("SELECT * FROM user WHERE email = ? AND provider = SELF", email)
	if error != nil {
		return nil, error
	}
	return &user, nil
}
