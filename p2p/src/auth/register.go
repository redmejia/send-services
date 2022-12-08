package auth

type Register struct {
	UserUID  string `json:"user_uid"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"pwd"`
	Bank     `json:"bank"`
}
