package login

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UID      string `json:"uid"`
	GranType string `json:"grant_type"`
}
