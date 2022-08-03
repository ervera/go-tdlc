package encrypt

import "golang.org/x/crypto/bcrypt"

func Password(pass string) (string, error) {
	cost := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
