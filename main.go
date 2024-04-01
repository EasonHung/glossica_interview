package main

import (
	"database/sql"
	"glossika_be_interview/controllers/recommendation_controller"
	"glossika_be_interview/controllers/user_controller"
	"glossika_be_interview/db_client"
	"glossika_be_interview/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

func Route(route *gin.Engine) {
	userRoute := route.Group("/user")
	{
		userRoute.POST("/create", user_controller.CreateNewUser)
		userRoute.POST("/login", user_controller.Login)
	}

	emailVerifyRoute := route.Group("/emailVerify")
	{
		emailVerifyRoute.POST("/send", user_controller.SendEmailVerification)
		emailVerifyRoute.POST("/verify", user_controller.VerifyEmail)
	}

	recommendationRoute := route.Group("/recommendation", middleware.VerifyToken())
	{
		recommendationRoute.GET("/all", recommendation_controller.GetRecommendations)
	}
}

func main() {
	// Connect to MySQL
	var err error
	db_client.DB, err = sql.Open("mysql", "user:password@tcp(glossika-mysql:3306)/glossika_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db_client.DB.Close()

	// Connect to Redis
	db_client.Rdb = redis.NewClient(&redis.Options{
		Addr: "glossika-redis:6379",
	})
	defer db_client.Rdb.Close()

	// Initialize Gin
	router := gin.Default()

	Route(router)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
