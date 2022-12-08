package auth

type Success struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}

type Fail struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
}
