package user_service

import "cn.jalivv.code/bytedance-douyin/models"

type User struct {
	models.Model
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
	Password      string
}

func CheckUserExist(u *models.User) (bool, error) {
	return u.CheckUserExist()
}

func Save(u *models.User) error {
	return u.Save()
}
