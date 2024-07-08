package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/leMoreira/api_simples/cmd/api/bd"
	"github.com/leMoreira/api_simples/cmd/api/models"
	"golang.org/x/crypto/bcrypt"
)

func Createuser(ctx *gin.Context) {
	var authInput models.AuthInput

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	var userFound models.User

	bd.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username já está sendo usado",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	bd.DB.Create(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func Login(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	bd.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte("testeApi"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}
