package auth

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"pwd"`
}
