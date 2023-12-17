package main

import (
	"dbo-technical-test/config"
	"dbo-technical-test/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := config.ConnectDB()
	route := routers.RouterConfig(db)

	route.Run(os.Getenv("APP_PORT"))
}
