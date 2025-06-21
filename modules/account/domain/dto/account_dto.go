package dto

type SignUpAccount struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAccount struct{}
