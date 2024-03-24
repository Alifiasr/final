package router

import (
	"MyGram/controller"
	"MyGram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.DELETE("/delete/:userId", controller.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/create", controller.CreatePhoto)
		photoRouter.GET("/getall", controller.GetAllPhotos)
		photoRouter.GET("/get/:photoId", controller.GetPhoto)
		photoRouter.PUT("/update/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/create/:photoId", controller.CreateComment)
		commentRouter.GET("/getall", controller.GetAllComments)
		commentRouter.GET("/getall/:photoId", controller.GetAllCommentsForPhoto)
		commentRouter.GET("/get/:commentId", controller.GetComment)
		commentRouter.PUT("/update/:commentId", middleware.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/delete/:commentId", middleware.CommentAuthorization(), controller.DeleteComment)
	}

	sosialMediaRouter := r.Group("/sosialmedia")
	{
		sosialMediaRouter.Use(middleware.Authentication())
		sosialMediaRouter.POST("/create", controller.CreateSocialMedia)
		sosialMediaRouter.GET("/getall", controller.GetAllSocialMedias)
		sosialMediaRouter.GET("/get/:socialMediaId", controller.GetSocialMedia)
		sosialMediaRouter.PUT("/update/:socialMediaId", middleware.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		sosialMediaRouter.DELETE("/delete/:socialMediaId", middleware.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}

	return r
}
