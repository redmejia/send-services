package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password *string) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	if err != nil {
		return err
	}
	*password = string(hashedPwd)

	return nil
}

func ComparePassword(hashedPwd string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
	// if err == nil {
	// 	return false
	// }

	return err == nil

}
