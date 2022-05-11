package video_service

import "cn.jalivv.code/bytedance-douyin/models"

func Feed(startTime string) (*[]models.Video, error) {
	v := models.Video{}
	return v.GetFeed()

}
