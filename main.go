package main

import (
	"github.com/ehilmidag/controllers"
	"github.com/ehilmidag/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run() //
}

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDatabase()
	utils.SyncDatabase()

}
