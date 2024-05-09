package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	app := gin.Default()

	PORT := os.Getenv("PORT")

	app.Run(":" + PORT)
}
