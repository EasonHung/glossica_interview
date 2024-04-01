package email_verify_service

import (
	"context"
	"fmt"
	"glossika_be_interview/db_client"
	"glossika_be_interview/domains/user/user_repository"
	"math/rand"
	"time"
)

func SendEmailVerification(userId string) (string, error) {
	code := generateRandomNumberString()

	err := db_client.Rdb.Set(context.TODO(), "email-verify-"+code, userId, time.Minute*time.Duration(10)).Err()
	if err != nil {
		return code, err
	}

	err = sendEmail()
	if err != nil {
		return code, err
	}
	return code, nil
}

func VerifyEmail(code string) error {
	userId, err := db_client.Rdb.Get(context.TODO(), "email-verify-"+code).Result()
	if err != nil {
		return err
	}

	user, err := user_repository.FindUserByUserId(userId)
	if err != nil {
		return err
	}

	err = user_repository.UpdateEmailStatus(user.UserId, true)
	if err != nil {
		return err
	}

	return nil
}

func sendEmail() error {
	return nil
}

func generateRandomNumberString() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random four-digit number
	randomNumber := rand.Intn(10000)
	// Ensure the number is exactly four digits
	randomNumber %= 10000

	// Convert the number to a string with leading zeros if necessary
	randomNumberString := fmt.Sprintf("%04d", randomNumber)

	return randomNumberString
}
