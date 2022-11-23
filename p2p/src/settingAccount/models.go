package settingaccount

type Bank struct {
	FullName string `json:"full_name"`
	Card     string `json:"card"`
	// more information here
}

type Register struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"pwd"`
	Bank     `json:"bank"`
}

type Signin struct {
	Email    string `json:"email"`
	Password string `json:"pwd"`
}
