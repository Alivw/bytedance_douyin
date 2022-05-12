package models

type UserLikeVideo struct {
	Model
	UserID     int64 `gorm:"type:bigint" json:"user_id"`
	VideoID    int64 `gorm:"type:bigint" json:"video_id"`
	ActionType int32 `json:"action_type" gorm:"-"`
}
