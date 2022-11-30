package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aryaw/urlshortner/middleware"
	"github.com/aryaw/urlshortner/config"
	"github.com/aryaw/urlshortner/src/user"
	"github.com/aryaw/urlshortner/src/authuser"
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

	v1 := r.Group("/v1")
	{
		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.GET("/user/logout", user.Logout)

		v1.POST("/token/refresh", authuser.Refresh)
	}
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	r.Run(":" + port)

	r.Routes()
}
