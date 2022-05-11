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

func (u *User) CheckUserExist() (bool, error) {
	err := db.Select("id").Where(u).First(u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("databese error")
	}
	if u.ID > 0 {
		return true, errors.New("User already exist")
	}
	return false, nil
}

func (u *User) Save() (int32, error) {
	if err := db.Save(u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
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
