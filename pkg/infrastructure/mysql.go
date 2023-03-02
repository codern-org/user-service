package infrastructure

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // Load MySQL driver
	"github.com/jmoiron/sqlx"
)

func OpenMySqlConnection(uri string) (*sqlx.DB, error) {
	db, error := sqlx.Connect("mysql", uri)

	if error != nil {
		return nil, error
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
