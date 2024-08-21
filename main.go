package main

import (
	"log"
	"net/http"

	"github.com/hazaloolu/blog-api/internal/router"
	"github.com/hazaloolu/blog-api/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	storage.InitDB()
	r := router.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))

}
