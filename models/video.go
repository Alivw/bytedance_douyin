package models

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
func (Video) TableName() string {
	return "t_video"
}
