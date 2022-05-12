package controller

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"cn.jalivv.code/bytedance-douyin/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteReq struct {
	UserID     int64  `form:"user_id"`
	Token      string `form:"token"`
	VideoID    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"`
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var f FavoriteReq
	c.Bind(&f)
	ulv := models.UserLikeVideo{
		UserID:     f.UserID,
		VideoID:    f.VideoID,
		ActionType: f.ActionType,
	}
	if err := user_service.FavoriteAction(&ulv); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		panic(err)
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
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
