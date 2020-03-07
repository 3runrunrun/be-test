package middleware

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// load .env file
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Println("middleware/main.go : ", err)
	}
	log.Println("middleware/main.go : .env file loaded for middleware")

	// call init function for each middlewares
	// we call init of jwt.go
	initJwtMiddleware()
	// add some below, if you need more
	// init...()
	// init...()

}

// get .env key
// probably will be used in accross middlewares
func getEnvKey(key string) string {
	ret, flag := os.LookupEnv(key)
	if !flag {
		log.Println("middleware/constantsetter.go : key not found")
		return ""
	}

	return ret
}
