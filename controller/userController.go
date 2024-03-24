package controller

import (
	"MyGram/config"
	"MyGram/helper"
	"MyGram/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	db := config.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"username": User.Username,
		"email":    User.Email,
		"age":      User.Age,
	})
}

func UserLogin(c *gin.Context) {
	db := config.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Username, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func DeleteUser(c *gin.Context) {
	db := config.GetDB()
	User := model.User{}

	userId, _ := strconv.Atoi(c.Param("userId"))

	err := db.Where("id = ?", userId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The photo has been successfully deleted",
	})
}
