package user_service

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"cn.jalivv.code/bytedance-douyin/pkg/gredis"
	"encoding/json"
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

	// hset user:userId  videi_id videoJson

	var video = models.Video{}
	err := (&video).GetByID()
	if err != nil {
		return err
	}
	// 点赞操作

	video.IsFavorite = true
	vbytes, err := json.Marshal(video)
	if err != nil {
		return err
	}
	if ulv.ActionType == 1 {

		err = gredis.SADD(fmt.Sprintf("%s%v", "user_like_video:", ulv.UserID), string(vbytes))
	} else {
		//video.IsFavorite = true
		//vbytes, err := json.Marshal(video)
		//if err != nil {
		//	return err
		//}
		err = gredis.Delete(fmt.Sprintf("%s%v", "user_like_video:", ulv.UserID), string(vbytes))
	}

	return err
}
