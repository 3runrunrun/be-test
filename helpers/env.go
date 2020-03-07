package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./../.env")
	if err != nil {
		log.Println("utils/env.go : ", err)
	}
}

// GetKey to get specific env keys
func GetKey(key string) string {
	ret, flag := os.LookupEnv(key)
	if !flag {
		log.Println("utils/env.go : cannot lookup to .env keys")
		return ""
	}

	return ret
}
