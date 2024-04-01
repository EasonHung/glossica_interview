package user_repository

import (
	"glossika_be_interview/db_client"
	"glossika_be_interview/domains/user/user_entities"
)

func CreateUser(email string, password string) error {
	newUser := user_entities.NewUser(email, password)

	_, err := db_client.DB.Exec("INSERT INTO users(userId, email, password) VALUES(?, ?, ?)", newUser.UserId, newUser.Email, newUser.Password)
	return err
}

func FindUserByUserId(userId string) (user_entities.User, error) {
	var user user_entities.User
	err := db_client.DB.QueryRow("SELECT userId, email, password FROM users WHERE userId = ?", userId).Scan(&user.UserId, &user.Email, &user.Password)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func FindUserByEmail(email string) (user_entities.User, error) {
	var user user_entities.User
	err := db_client.DB.QueryRow("SELECT userId, email, password FROM users WHERE email = ?", email).Scan(&user.UserId, &user.Email, &user.Password)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func UpdateEmailStatus(userId string, emailStatus bool) error {
	_, err := db_client.DB.Exec("UPDATE users SET verfiedEmail = ? WHERE userId = ?", emailStatus, userId)
    if err != nil {
        return err
    }

	return nil
}
