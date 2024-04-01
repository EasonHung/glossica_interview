package user_controller

import (
	"glossika_be_interview/dto"
	"glossika_be_interview/services/email_verify_service"
	"glossika_be_interview/services/user_service"

	"github.com/gin-gonic/gin"
)

func CreateNewUser(c *gin.Context) {
	var request dto.CreateNewUserReq
	c.BindJSON(&request)

	err := user_service.CreateUser(request.Email, request.Password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
	return
}

func Login(c *gin.Context) {
	var request dto.LoginReq
	c.BindJSON(&request)

	token, err := user_service.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, token)
	return
}

func SendEmailVerification(c *gin.Context) {
	var request dto.SendEmailVerificationReq
	c.BindJSON(&request)

	code, err := email_verify_service.SendEmailVerification(request.UserId)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, code)
	return
}

func VerifyEmail(c *gin.Context) {
	var request dto.VerifyEmailReq
	c.BindJSON(&request)

	err := email_verify_service.VerifyEmail(request.Code)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
	return
}
