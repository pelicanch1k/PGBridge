package router

import (
	"PGBridge/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(c *controller.Controller) *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	get := router.Group("/api/v1")
	{
		get.GET("/user/:id", c.GetUserByID)
		get.GET("/comment/:id", c.GetCommentByID)

		get.GET("user/:id/comment", c.GetListComment)
		get.GET("user/:id/comment/:id", c.GetCommentByUserAndID)
	}

	post := router.Group("/api/v1")
	{
		post.POST("/user/:id/comment", c.CreateComment)
		post.POST("/user/", c.CreateUser)
	}

	upd := router.Group("/api/v1")
	{
		upd.PUT("/user/:id", c.UpdateUser)
		upd.PUT("/comment/:id", c.UpdateComment)
	}

	del := router.Group("/api/v1")
	{
		del.DELETE("/comment/:id", c.DeleteComment)
		del.DELETE("/user/:id", c.DeleteUser)
	}

	return router
}
