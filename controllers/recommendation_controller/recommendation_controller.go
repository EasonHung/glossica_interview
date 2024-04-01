package recommendation_controller

import (
	"glossika_be_interview/services/recommendation_service"

	"github.com/gin-gonic/gin"
)

func GetRecommendations(c *gin.Context) {
	recommendations, err := recommendation_service.GetAll()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, recommendations)
	return
}
