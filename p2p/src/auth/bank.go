package auth

type Bank struct {
	FullName string `json:"full_name"`
	Card     string `json:"card"`
	CvNumber string `json:"cv_number"`
	// more information here
}
