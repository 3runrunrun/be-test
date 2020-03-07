package middleware

import (
	"log"
	"net/http"

	"github.com/3runrunrun/be-test/helpers"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authorization with JWT
func Authorization() gin.HandlerFunc {

	// get .env key for JWT setting
	appkey := helpers.GetKey("APPKEY")

	return func(c *gin.Context) {

		// create jwtMiddelWare instance
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(appkey), nil
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
