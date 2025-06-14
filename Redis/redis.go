package Redis

//import (
//	"context"
//	"github.com/LJin-go/exam/global"
//	"github.com/go-redis/redis/v8"
//	"log"
//)
//
//func Redis(addr, password string, db int) {
//	global.RDB = redis.NewClient(&redis.Options{
//		Addr:     addr,
//		Password: password,
//		DB:       db,
//	})
//
//	err := global.RDB.Ping(context.Background()).Err()
//	if err != nil {
//		panic("redis连接失败")
//	}
//	log.Println("redis连接成功")
//}
