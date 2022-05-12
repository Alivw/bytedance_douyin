package user_service

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"cn.jalivv.code/bytedance-douyin/pkg/gredis"
	"fmt"
)

type User struct {
	models.Model
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
	Password      string
}

func CheckUserExist(u *models.User) (bool, error) {
	return u.CheckUserExists()
}

func Save(u *models.User) error {
	return u.Save()
}

func FavoriteAction(ulv *models.UserLikeVideo) error {
	// 点赞操作
	var err error
	if ulv.ActionType == 1 {
		err = gredis.SAdd(fmt.Sprintf("%s%v", "user_like_video:", ulv.UserID), ulv.VideoID)
	} else {
		err = gredis.SREM(fmt.Sprintf("%s%v", "user_like_video:", ulv.UserID), ulv.VideoID)
	}

	return err
}
