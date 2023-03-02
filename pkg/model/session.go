package model

type Session struct {
	Id        string `json:"id" db:"id"`
	UserId    string `json:"userId" db:"user_id"`
	IpAddress string `json:"ipAddress" db:"ip_address"`
	UserAgent string `json:"userAgent" db:"user_agent"`
	ExpiredAt int64  `json:"expiredAt" db:"expired_at"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
}
