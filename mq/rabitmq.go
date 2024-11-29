package mq

import (
	"github.com/streadway/amqp"
	"walletserver/db"
	"walletserver/log"
)

var gMQCh *amqp.Channel
var gQueue *amqp.Queue
var gDeliveries <-chan amqp.Delivery

var gMQName = "game_bet_history"

// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
func InitRabitmq() {

	//fmt.Println(GetConfig().RabitMQ.Host)
	conn, err := amqp.Dial(db.GetConfig().RabitMQ.Host)
	if err != nil {
		log.Error("无法连接到 mq! err:", err, db.GetConfig().RabitMQ.Host)
		panic(err)
	}
	//defer conn.Close()

	CreateMQChannel(conn)
}

// 创建通道（Channel）
func CreateMQChannel(conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("创建通道失败!err:", err)
		panic(err)
	}
	//defer ch.Close()
	gMQCh = ch

	CreateMQQueue()
}

// 创建消息队列
func CreateMQQueue() {
	queue, err := gMQCh.QueueDeclare(
		gMQName, // queue name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal("创建消息队列失败!err:", err)
		panic(err)
	} else {
		gQueue = &queue
	}
}

// 发送
func PublishMQMsg(msg []byte) {
	err := gMQCh.Publish(
		"",          // exchange
		gQueue.Name, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	if err != nil {
		log.Error("发送消息失败!err:", err, " msg:", string(msg))
	}
}

// 消费者
func CreateComsume() {
	deliveries, err := gMQCh.Consume(
		gQueue.Name, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Fatal("创建消费者失败!err:", err)
		panic(err)
	} else {
		gDeliveries = deliveries
	}
}

func GetRabitMQClient() (*amqp.Channel, *amqp.Queue, <-chan amqp.Delivery) {
	return gMQCh, gQueue, gDeliveries
}
