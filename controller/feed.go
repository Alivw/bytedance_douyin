package controller

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"cn.jalivv.code/bytedance-douyin/service/video_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList *[]models.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	query := c.DefaultQuery("latest_time", strconv.FormatInt(time.Now().Unix(), 10))
	t, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		panic(err)
	}
	strTime := time.Unix(t, 0).Format("2006-01-02 15:04:05")

	videos, err := video_service.Feed(strTime)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1}})
		panic(err)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
func VidelFile(c *gin.Context) {
	name := c.Param("name")
	fileName := path.Join("./public/", name)
	c.File(fileName)
}
