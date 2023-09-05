package controllers

import (
	"github.com/ehilmidag/models"
	"github.com/ehilmidag/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Signup(ctx *gin.Context) {
	var body *models.SignUp

	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := models.User{
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	result := utils.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	ctx.JSON(200, gin.H{})
}
