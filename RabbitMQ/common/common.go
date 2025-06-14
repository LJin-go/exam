package common

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/LJin-go/exam/global"
//	"github.com/streadway/amqp"
//	"log"
//)
//
//func QueuePub[T any](queue string, data T) {
//
//	message, err := json.Marshal(data)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
//	_, err = global.MQ.QueueDeclare(
//		queue,
//		//是否持久化
//		false,
//		//是否自动删除
//		false,
//		//是否具有排他性
//		false,
//		//是否阻塞处理
//		false,
//		//额外的属性
//		nil,
//	)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	//调用channel 发送消息到队列中
//	_ = global.MQ.Publish(
//		"",
//		queue,
//		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
//		false,
//		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
//		false,
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        message,
//		})
//}
//
//func QueueSub(queue string) <-chan amqp.Delivery {
//	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
//	q, err := global.MQ.QueueDeclare(
//		queue,
//		//是否持久化
//		false,
//		//是否自动删除
//		false,
//		//是否具有排他性
//		false,
//		//是否阻塞处理
//		false,
//		//额外的属性
//		nil,
//	)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	//接收消息
//	result, err := global.MQ.Consume(
//		q.Name, // queue
//		//用来区分多个消费者
//		"", // consumer
//		//是否自动应答
//		true, // auto-ack
//		//是否独有
//		false, // exclusive
//		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
//		false, // no-local
//		//列是否阻塞
//		false, // no-wait
//		nil,   // args
//	)
//
//	return result
//}
//
//func PublishPub[T any](exchange string, data T) {
//
//	message, _ := json.Marshal(data)
//
//	//1.尝试创建交换机
//	_ = global.MQ.ExchangeDeclare(
//		exchange,
//		"fanout",
//		true,
//		false,
//		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
//		false,
//		false,
//		nil,
//	)
//
//	//2.发送消息
//	_ = global.MQ.Publish(
//		exchange,
//		"",
//		false,
//		false,
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        message,
//		})
//}
//
//func PublishSub(exchange string) <-chan amqp.Delivery {
//
//	_ = global.MQ.ExchangeDeclare(
//		exchange,
//		//交换机类型
//		"fanout",
//		true,
//		false,
//		//YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
//		false,
//		false,
//		nil,
//	)
//
//	//2.试探性创建队列，这里注意队列名称不要写
//	q, _ := global.MQ.QueueDeclare(
//		"", //随机生产队列名称
//		false,
//		false,
//		true,
//		false,
//		nil,
//	)
//
//	//绑定队列到 exchange 中
//	_ = global.MQ.QueueBind(
//		q.Name,
//		//在pub/sub模式下，这里的key要为空
//		"",
//		exchange,
//		false,
//		nil)
//
//	//消费消息
//	result, _ := global.MQ.Consume(
//		q.Name,
//		"",
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//
//	return result
//}
