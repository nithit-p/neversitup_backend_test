package domain

type Auth struct {
	AuthId       int    `json:"auth_id"`
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
