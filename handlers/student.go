package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
}
