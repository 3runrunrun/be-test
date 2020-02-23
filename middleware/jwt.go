package middleware

import (
	"log"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// GetCredential information from .env
func GetCredential() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Println("middleware jwt.go: ", err)
			c.Next()
		}

		v, flag := os.LookupEnv("APPKEY")
		if !flag {
			log.Println("middleware jwt.go: ", flag)
			c.Next()
		}

		log.Println(v)
		c.Set("appkey", v)
		c.Next()
	}
}

// Authorization with JWT
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("smartlink"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		err := jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			log.Println("middleware jwt.go: ", err)
			c.Writer.Header().Set("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
