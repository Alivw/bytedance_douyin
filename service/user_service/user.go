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

func (u *User) CheckUserExist() (bool, error) {
	user := &models.User{Name: u.Name, Password: u.Password}
	return user.CheckUserExist()
}

func (u *User) Save() (int32, error) {
	m := &models.User{Name: u.Name, Password: u.Password}
	return m.Save()
}
