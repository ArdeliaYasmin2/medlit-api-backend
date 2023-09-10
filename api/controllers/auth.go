package controllers

import (
	"fmt"
	"net/http"

	"medlit-api-backend/api/models"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *ControllerRepo) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password
	user.Name = u.Repo.GetNameByEmail(user.Email)

	token, err := u.Repo.LoginCheck(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   "false",
		"message": "Login Successful",
		"loginResult": gin.H{
			"name":  user.Name,
			"token": token,
		},
	})
}

type RegisterInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func (u *ControllerRepo) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	user := models.User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	errConfirmPass := u.Repo.RegisterCheck(input.Password, input.ConfirmPassword)
	if errConfirmPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": errConfirmPass.Error(),
		})
		return
	}

	result, err := u.Repo.SaveUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	fmt.Println(result)

	c.JSON(http.StatusOK, gin.H{
		"error":   "false",
		"message": "Registration Successful",
	})
}
