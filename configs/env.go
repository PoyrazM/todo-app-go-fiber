package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Env file was not load")
	}

	mongoURI := os.Getenv("MONGO_URI")
	return mongoURI
}
