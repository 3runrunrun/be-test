package middleware

import (
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtAppKey string
var jwtSigntureKey string

// initializing .env loader for middleware only
func initJwtMiddleware() {
	setJwtAppKey()       // set jwt appkey
	setJwtSignatureKey() // set jwt signature_key
}

// set constant jwtAppkey
func setJwtAppKey() {
	key := &jwtAppKey
	*key = getEnvKey("APPKEY")
}

// set constant jwt
func setJwtSignatureKey() {
	key := &jwtSigntureKey
	*key = getEnvKey("SIGNATUREKEY")
}

// Authorization with JWT
func Authorization() gin.HandlerFunc {

	return func(c *gin.Context) {

		// create jwtMiddleWare instance
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtAppKey), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		// check JWT along with incoming request
		// remember! JWT written in request Header
		err := jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			log.Println("middleware/jwt.go: ", err)

			c.Writer.Header().Set("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "unauthorized"})
			c.Abort()

			return
		}
		c.Next()
	}
}
