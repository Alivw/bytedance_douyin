package routers

import (
	"cn.jalivv.code/bytedance-douyin/controller"
	"cn.jalivv.code/bytedance-douyin/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// public directory is used to serve static resources
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.Use(jwt.JWT())
	{
		apiRouter.POST("/publish/action/", controller.PublishWithoutOss)
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)
		apiRouter.GET("/comment/list/", controller.CommentList)

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	}

	//apiRouter.GET("/public/:name", controller.VidelFile)

	return r
}
