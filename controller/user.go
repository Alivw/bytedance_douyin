package controller

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"cn.jalivv.code/bytedance-douyin/pkg/util"
	"cn.jalivv.code/bytedance-douyin/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User models.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := &models.User{Name: util.EncodeMD5(username), Password: util.EncodeMD5(password)}
	//	Determine whether the user exists.
	if exist, _ := user_service.CheckUserExist(user); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		// register
		err := user_service.Save(user)
		if err != nil {

			panic(err)
		}

		token, err := util.GenerateToken(username, password)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   int64(user.ID),
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	m := models.User{Name: util.EncodeMD5(username), Password: util.EncodeMD5(password)}
	exist, err := user_service.CheckUserExist(&m)
	if !exist || err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		token, _ := util.GenerateToken(username, password)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   int64(m.ID),
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0, StatusMsg: "success"},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}

	claims, _ := util.ParseToken(c.Query("token"))

	user := models.User{Name: claims.Username, Password: claims.Password}
	exist, err := user_service.CheckUserExist(&user)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		panic(err)
	}
	if !exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用不不存在"},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0, StatusMsg: "success"},
		User:     user,
	})

}
