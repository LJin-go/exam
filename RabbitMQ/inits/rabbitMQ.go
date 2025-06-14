package inits

//import (
//	"fmt"
//	"github.com/LJin-go/exam/global"
//	"github.com/streadway/amqp"
//)
//
//func InitRabbitMq() {
//	//获取connection
//	conn, err := amqp.Dial("amqp://username:password@127.0.0.1:5672/vhost")
//	if err != nil {
//		panic("Failed to connect to RabbitMQ")
//	}
//	//获取channel
//	global.MQ, err = conn.Channel()
//
//	if err != nil {
//		panic("Failed to open a channel")
//	}
//
//	fmt.Println("Connected to RabbitMQ")
//
//}
