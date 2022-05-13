package gredis

import (
	"cn.jalivv.code/bytedance-douyin/pkg/setting"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

//var RedisConn *redis.Pool
var ctx = context.Background()

var rdb *redis.Client

// Setup Initialize the Redis instance
//func Setup() error {
//	RedisConn = &redis.Pool{
//		MaxIdle:     setting.RedisSetting.MaxIdle,
//		MaxActive:   setting.RedisSetting.MaxActive,
//		IdleTimeout: setting.RedisSetting.IdleTimeout,
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
//			if err != nil {
//				return nil, err
//			}
//			if setting.RedisSetting.Password != "" {
//				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
//					c.Close()
//					return nil, err
//				}
//			}
//			return c, err
//		},
//		TestOnBorrow: func(c redis.Conn, t time.Time) error {
//			_, err := c.Do("PING")
//			return err
//		},
//	}
//
//	return nil
//}
func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.RedisSetting.Host,
		Password: setting.RedisSetting.Password, // no password set
		DB:       0,                             // use default DB
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

}

// Set a key/value
//func Set(key string, data interface{}, time int) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	val, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	_, err = conn.Do("SET", key, val)
//	if err != nil {
//		return err
//	}
//	if time != 0 {
//		_, err = conn.Do("EXPIRE", key, time)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

//
//// Exists check a key
//func Exists(key string) bool {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	exists, err := redis.Bool(conn.Do("EXISTS", key))
//	if err != nil {
//		return false
//	}
//
//	return exists
//}
//
//// Get get a key
//func Get(key string) ([]byte, error) {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	reply, err := redis.Bytes(conn.Do("GET", key))
//	if err != nil {
//		return nil, err
//	}
//
//	return reply, nil
//}
//
//// Delete delete a kye
//func Delete(key string) (bool, error) {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	return redis.Bool(conn.Do("DEL", key))
//}
//
//// LikeDeletes batch delete
//func LikeDeletes(key string) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//
//	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
//	if err != nil {
//		return err
//	}
//
//	for _, key := range keys {
//		_, err = Delete(key)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func SAdd(key string, val interface{}) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//	_, err := conn.Do("SADD", key, val)
//	return err
//}
//
//func SREM(key string, val interface{}) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//	_, err := conn.Do("SREM", key, val)
//	return err
//}
//
//func Option(opt string, key string, val interface{}) error {
//	conn := RedisConn.Get()
//	defer conn.Close()
//	_, err := conn.Do(opt, key, val)
//	return err
//}

func SADD(key string, val string) error {
	_, err := rdb.SAdd(ctx, key, val).Result()
	return err
}

func Delete(key string, member string) error {
	rdb.SRem(ctx, key, member)
	return nil
}
