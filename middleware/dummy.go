package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Dummy middleware
func Dummy() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("this is dummy middleware")
		c.Next()
	}
}
