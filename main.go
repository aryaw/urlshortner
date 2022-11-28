package main

import (
	"log"
	"os"

	"urlshortner/middleware"
	"urlshortner/config"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// database
	DB := config.Init()
    h := config.New(DB)
	
	config.InitRedis(1)
}
