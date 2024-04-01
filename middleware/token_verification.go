package middleware

import (
	"glossika_be_interview/services/token_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.GetHeader("UserToken")

		err, _ := token_service.VerifyAccessToken(userToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Process request
		c.Next()
	}
}
