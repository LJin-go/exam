package main

//import (
//	"github.com/LJin-go/exam/RabbitMQ/common"
//	"github.com/LJin-go/exam/RabbitMQ/inits"
//	"log"
//)
//
//func main() {
//	inits.InitRabbitMq()
//
//	msgS := common.PublishSub("publishTest")
//
//	forever := make(chan bool)
//	//启用协程处理消息
//	go func() {
//		for d := range msgS {
//			//消息逻辑处理，可以自行设计逻辑
//			log.Printf("Received a message: %s", d.Body)
//
//		}
//	}()
//
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//	<-forever
//}
