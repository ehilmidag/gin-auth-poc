package main

import (
	"github.com/ehilmidag/controllers"
	"github.com/ehilmidag/middleware"
	"github.com/ehilmidag/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequiredAuth, controllers.ValidateToken)
	r.Run() //
}

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDatabase()
	utils.SyncDatabase()

}
