package models

import (
	"encoding/json"
	"errors"
)

type Video struct {
	Model
	PlayUrl       string `gorm:"type:varchar(255);not null" json:"play_url"`
	CoverUrl      string `gorm:"type:varchar(255);not null" json:"cover_url"`
	FavoriteCount int64  `gorm:"type:int" json:"favorite_count"`
	CommentCount  int64  `gorm:"type:int" json:"comment_count"`
	IsFavorite    bool   `gorm:"type:tinyint(1)" json:"is_favorite"`
	UserID        uint   `gorm:"type:bigint " json:"user_id"`
	User          User   `json:"author";gorm:"foreignKey:UserID"`
}

//type BaseModel struct {
//	ID        uint       `gorm:"primarykey" json:"id"`
//	CreatedAt time.Time  `json:"created_at"`
//	UpdatedAt time.Time  `json:"updated_at"`
//	DeletedAt *time.Time `json:"deleted_at"`
//}

// TableName 会将 User 的表名重写为 `t_video`
//func (Video) TableName() string {
//	return "t_video"
//}

func (v Video) GetFeed() (*[]Video, error) {
	var vs = []Video{}
	// TODO 这里测试版本关闭 否则将导致获取不到比当前时间晚的视频
	//if err := repositoty.DB.Debug().Preload("User").Where("created_at>?", lastTime).Order("created_at desc").Find(&vs).Error; err != nil {
	//	panic(err)
	//}

	if err := db.Debug().Preload("User").Order("created_on desc").Find(&vs).Error; err != nil {
		return nil, err
	}
	return &vs, nil
}

func (v *Video) SaveFile(c chan string) error {
	url := <-c
	v.PlayUrl = url
	if err := db.Save(v).Error; err != nil {
		return err
	}

	return nil
}

func (v *Video) GetByID() error {
	db.First(v)
	if v.ID <= 0 {
		return errors.New("The liked video does not exist.")
	}
	return nil
}

func (v Video) MarshalBinary() (data []byte, err error) {
	return json.Marshal(v)
}

func (s Video) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}
