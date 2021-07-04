package api

import (
	"fmt"
	"log"
	"os"

	"github.com/ZootHii/blog-go-backend/api/controllers"
	"github.com/ZootHii/blog-go-backend/api/utils/seeds"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env %v", err)
	} else {
		fmt.Println("Getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seeds.Load(server.DB)

	server.Run(":8010")

}
