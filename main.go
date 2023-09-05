package main

import (
	"github.com/ehilmidag/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() //
}

func init() {
	var DB utils.DB
	utils.LoadEnvVariables()
	DB.ConnectToDatabase()
}
