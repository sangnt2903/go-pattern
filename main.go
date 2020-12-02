package main

import (
	"github.com/gin-gonic/gin"
	"go-pattern/config"
)

func main() {
	//connect database
	config.ConnectToDatabase()
	// router
	r := gin.Default()
	config.Run(r, "127.0.0.1", "8089")
}
