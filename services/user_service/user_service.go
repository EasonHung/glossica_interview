package user_service

import (
	"errors"
	"glossika_be_interview/domains/user/user_repository"
	"glossika_be_interview/services/token_service"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(email string, password string) error {
	if !validatePassword(password) {
		return errors.New("password invalid")
	}

	encryptedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return user_repository.CreateUser(email, string(encryptedPasswordBytes))
}

func Login(email string, password string) (string, error) {
	user, err := user_repository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	err, token := token_service.CreateAccessToken(user.UserId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func validatePassword(password string) bool {
	// Check length (between 6 and 16 characters)
	if len(password) < 6 || len(password) > 16 {
		return false
	}

	// Check for at least one uppercase letter
	upperRegex := regexp.MustCompile("[A-Z]")
	if !upperRegex.MatchString(password) {
		return false
	}

	// Check for at least one lowercase letter
	lowerRegex := regexp.MustCompile("[a-z]")
	if !lowerRegex.MatchString(password) {
		return false
	}

	// Check for at least one special character
	specialRegex := regexp.MustCompile(`[!@#$%^&*()_+\[\]{}<>+\-*/?,.:;"'_\\|~` + "`" + `!@#$%^&=]`)
	if !specialRegex.MatchString(password) {
		return false
	}

	// All conditions met
	return true
}
