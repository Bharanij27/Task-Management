package errorHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err string, status int) {
	message := fmt.Errorf(err)
	fmt.Println(message)
	c.JSON(status, gin.H{"error": message})
}
