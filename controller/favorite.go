package controller

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},

		/**
		var DemoVideos = []Video{
			{
				Id:            1,
				Author:        DemoUser,
				PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
				CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
			},
		}
		*/
		// TODO 为了没有完成这个接口导致app闪退，先把返回结果写死
		VideoList: []models.Video{
			{
				Model:         models.Model{ID: 1},
				PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
				CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
				UserID:        15,
				User: models.User{
					Model: models.Model{ID: 1},
					Name:  "jalivv",
				},
			},
		},
	})
}
