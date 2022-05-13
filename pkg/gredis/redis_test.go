package gredis

import (
	"cn.jalivv.code/bytedance-douyin/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestSet(t *testing.T) {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "service.aliww.top:6379",
		Password: "zhiruwakuangchengxumeiyoumasaidbyjalivv", // no password set
		DB:       0,                                         // use default DB
	})

	// 多个key-value值：HMSet、HMGet
	//m := make(map[string]interface{})
	//m["age"] = 18
	//m["address"] = "Japan"

	v := models.Video{
		UserID:  32,
		PlayUrl: "http://cn.pronhub.com",
	}
	marshal, _ := json.Marshal(v)
	_, err := rdb.LPush(ctx, "user", string(marshal)).Result()
	if err != nil {
		panic(err)
	}
	//if _, err := rdb.HMSet(ctx, "user:15", 1, 2).Result(); err != nil {
	//	fmt.Printf("%v", err)
	//}
}
