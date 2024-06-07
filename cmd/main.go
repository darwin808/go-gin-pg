package main

import (
	"log"
	"time"

	"github.com/darwin808/pg-gin/database"
	"github.com/darwin808/pg-gin/handlers"
	"github.com/gin-contrib/cors"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDb()

	r := gin.New()

	logger, _ := zap.NewProduction()

	r.Use(cors.Default())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	r.Use(ginzap.RecoveryWithZap(logger, true))

	v1 := r.Group("/v1")
	{
		v1.GET("/", handlers.Home)
		v1.GET("/students", handlers.GetStudents)
		v1.POST("/students", handlers.CreateStudent)
		v1.GET("/students/:id", handlers.FindStudent)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run(":3000")
}
