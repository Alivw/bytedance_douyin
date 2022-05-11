package controller

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"cn.jalivv.code/bytedance-douyin/pkg/setting"
	"cn.jalivv.code/bytedance-douyin/pkg/util"
	"cn.jalivv.code/bytedance-douyin/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
)

type VideoListResponse struct {
	Response
	VideoList []models.Video `json:"video_list"`
}

//Publish check token then save upload file to public directory
//func Publish(c *gin.Context) {
//
//	token := c.Query("token")
//	// TODO 登陆权限校验
//	//if _, exist := usersLoginInfo[token]; !exist {
//	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	//	return
//	//}
//	// 从请求中获取文件
//	data, err := c.FormFile("data")
//	if err != nil {
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//	filename := filepath.Base(data.Filename)
//	user := usersLoginInfo[token]
//	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
//	fileHandle, err := data.Open()
//	if err != nil {
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//	fileByte, err := ioutil.ReadAll(fileHandle) //获取上传文件字节流
//	if err != nil {
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//	defer fileHandle.Close()
//
//	var video = models.Video{UserID: uint(user.Id)}
//	// 上传文件
//	service.NewVideoServiceInstance().UploadFile(fileByte, finalName, &video)
//
//	c.JSON(http.StatusOK, Response{
//		StatusCode: 0,
//		StatusMsg:  finalName + " uploaded successfully",
//	})
//}

//PublishWithoutOss The video uploaded by the user is saved to the local server.
func PublishWithoutOss(c *gin.Context) {
	token := c.PostForm("token")
	claims, err := util.ParseToken(token)
	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	u := models.User{Name: claims.Username, Password: util.EncodeMD5(claims.Password)}
	user_service.CheckUserExist(&u)
	finalName := fmt.Sprintf("%d_%s", u.ID, filename)
	saveFile := filepath.Join("./public/", finalName)

	var sc = make(chan string, 1)
	group := sync.WaitGroup{}
	group.Add(2)
	// do save file to disk
	go func() {
		defer close(sc)
		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
		host := setting.VideoUriSetting.Host
		if !strings.HasSuffix(host, "/") {
			host = fmt.Sprintf("%s%s", host, "/")
		}
		sc <- fmt.Sprintf("%s%s%s", host, "static/", finalName)
		group.Done()
	}()

	// do Save data to database
	go func() {
		video := &models.Video{UserID: uint(u.ID)}
		video.SaveFile(sc)
		group.Done()
	}()

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	claims, _ := util.ParseToken(c.Query("token"))

	user := models.User{Name: claims.Username, Password: claims.Password}
	user_service.CheckUserExist(&user)
	videoList, err := user.GetPublishList()
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
			},
		})
		panic(err)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
