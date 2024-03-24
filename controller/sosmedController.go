package controller

import (
	"MyGram/config"
	"MyGram/helper"
	"MyGram/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	SocialMedia := model.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetAllSocialMedias(c *gin.Context) {
	db := config.GetDB()
	allSocialMedias := []model.SocialMedia{}

	db.Find(&allSocialMedias)

	if len(allSocialMedias) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No social media found",
			"error_message": "There are no social media found.",
		})
		return
	}

	c.JSON(http.StatusOK, allSocialMedias)
}

func GetSocialMedia(c *gin.Context) {
	db := config.GetDB()
	SocialMedia := model.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	SocialMedia.ID = uint(socialMediaId)

	err := db.First(&SocialMedia, "id = ?", socialMediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func UpdateSocialMedia(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	SocialMedia := model.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(model.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := config.GetDB()
	SocialMedia := model.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The social media has been successfully deleted",
	})
}
