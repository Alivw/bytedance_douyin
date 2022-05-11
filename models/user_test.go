package models

import (
	"cn.jalivv.code/bytedance-douyin/pkg/setting"
	"fmt"
	"testing"
)

func TestUser_GetPublishList(t *testing.T) {
	//vs := []Video{}
	//db.Where("user_id", u.ID).Find(&vs)
	//return vs, nil
	setting.Setup()
	Setup()
	u := &User{Model: Model{ID: 1}}
	vs, _ := u.GetPublishList()

	for _, v := range vs {
		fmt.Printf("%v", v)
	}

}
