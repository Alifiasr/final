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

func CreateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	contentType := helper.GetContentType(c)

	Comment := model.Comment{}
	userID := uint(userData["id"].(float64))

	Photo := model.Photo{}

	err := db.Select("user_id").First(&Photo, uint(photoId)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Photo Not Found",
			"message": "Photo doesn't exist, failed to create comment",
		})
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoId)

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetAllComments(c *gin.Context) {
	db := config.GetDB()
	allComments := []model.Comment{}

	db.Find(&allComments)

	if len(allComments) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No comments found",
			"error_message": "There are no comments found.",
		})
		return
	}

	c.JSON(http.StatusOK, allComments)
}

func GetAllCommentsForPhoto(c *gin.Context) {
	db := config.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	allComments := []model.Comment{}

	db.Where("photo_id = ?", photoId).Find(&allComments)

	if len(allComments) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No Comments found",
			"error_message": "There are no comments found for this photo.",
		})
		return
	}

	c.JSON(http.StatusOK, allComments)
}

func GetComment(c *gin.Context) {
	db := config.GetDB()
	Comment := model.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment.ID = uint(commentId)

	err := db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func UpdateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	Comment := model.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(model.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	db := config.GetDB()
	Comment := model.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The comment has been successfully deleted",
	})
}
