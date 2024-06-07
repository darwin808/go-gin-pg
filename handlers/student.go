package handlers

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/bytedance/sonic"
	"github.com/darwin808/pg-gin/database"
	"github.com/darwin808/pg-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Home(c *gin.Context) {
	c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
}

func GetStudents(c *gin.Context) {
	students := []models.Student{}
	// use Unscoped to remove deleted_at is null in query
	err := database.DB.Db.Unscoped().Find(&students).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"data": students})
}

func CreateStudent(c *gin.Context) {
	student := new(models.Student)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err})

		return
	}

	if err := sonic.Unmarshal(body, student); err != nil {
		c.JSON(500, gin.H{"error": err})

		return
	}

	// Check if email already exists
	var existingStudent models.Student
	if err := database.DB.Db.Where("email = ?", student.Email).First(&existingStudent).Error; err == nil {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err1 := database.DB.Db.Create(&student).Error
	if err1 != nil {
		c.JSON(500, gin.H{"error": err1})
		return
	}

	c.JSON(200, gin.H{"data": student})
}

func FindStudent(c *gin.Context) {
	student := models.Student{}
	id := c.Param("id")
	err := database.DB.Db.Find(&student, "id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err})
		return
	}

	if student.Email == "" {
		c.JSON(404, gin.H{"error": err})
		return
	}

	if student.Username == "" {
		c.JSON(404, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"data": student})
}
