package model

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GoogleUserResponse struct {
	Id              string `json:"sub"`
	Name            string `json:"name"`
	GivenName       string `json:"given_name"`
	FamilyName      string `json:"family_name"`
	ProfileUrl      string `json:"picture"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"email_verified"`
	Locale          string `json:"locale"`
}
