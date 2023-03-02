package model

type AuthProvider string

const (
	SELF   AuthProvider = "SELF"
	GOOGLE AuthProvider = "GOOGLE"
)

type User struct {
	Id          string       `json:"id" db:"id"`
	Email       string       `json:"email" db:"email"`
	Password    string       `json:"-" db:"password"`
	DisplayName string       `json:"displayName" db:"display_name"`
	ProfilePath string       `json:"profilePath" db:"profile_path"`
	Provider    AuthProvider `json:"provider" db:"provider"`
	CreatedAt   int64        `json:"createdAt" db:"created_at"`
}
