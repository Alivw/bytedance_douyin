package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	FollowCount   int64  `gorm:"type:int" json:"follow_count"`
	FollowerCount int64  `gorm:"type:int" json:"follower_count"`
	IsFollow      bool   `gorm:"type:tinyint(1)" json:"is_follow"`
	Password      string `gorm:"type:varchar(255);not null" json:"password"`
}

// CheckUserExists 判断用户是否存在，并且返回时会携带用户id
func (u *User) CheckUserExists() (bool, error) {
	err := db.Select("id").Where(u).First(u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("databese error")
	}
	if u.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (u *User) Save() error {
	if err := db.Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) GetPublishList() ([]Video, error) {
	var vs []Video
	//db.Debug().Where("user_id=?", u.ID).Find(&vs)
	db.Debug().Preload("User").Where("user_id=?", u.ID).Find(&vs)
	return vs, nil
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (int32, error) {
	var user User
	err := db.Select("id").Where(User{Name: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	if user.ID > 0 {
		return user.ID, nil
	}

	return 0, errors.New("用户不存在")
}
