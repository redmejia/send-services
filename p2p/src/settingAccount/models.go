package settingaccount

type Bank struct {
	FullName string `json:"full_name"`
	Card     string `json:"card"`
	CvNumber string `json:"cv_number"`
	// more information here
}

type Register struct {
	UserUID  string `json:"user_uid"`
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
