package Mysql

//import (
//	"fmt"
//	"github.com/LJin-go/exam/global"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
//func Mysql(user, pwd, host, db string, port int) {
//	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port, db)
//	var err error
//	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("MySQL连接失败")
//	}
//	log.Println("MySQL连接成功")
//}
