package user_entities

import "github.com/google/uuid"

type User struct {
	UserId       string `json:"userId"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	VerfiedEmail bool   `json:"verifiedEmail"`
}

func NewUser(email string, password string) User {
	userId := uuid.New()

	return User{
		UserId:       userId.String(),
		Email:        email,
		Password:     password,
		VerfiedEmail: false,
	}
}
