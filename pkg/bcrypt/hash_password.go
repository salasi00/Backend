package bcrypt

import "golang.org/x/crypto/bcrypt"

func HasingPassword(password string) (string, error) {
	hasedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hasedByte), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
