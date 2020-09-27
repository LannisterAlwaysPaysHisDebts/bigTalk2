/*
rabbitmq
 */
package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url格式 amqb://账号:密码@rabbitmq服务器地址:端口号/vhost名称
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 连接信息
	Mqurl string
}

// 实例化RabbitMQ
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbit := RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}

	var err error
	rabbit.conn, err = amqp.Dial(rabbit.Mqurl)
	rabbit.failOnErr(err, "创建连接错误")
	rabbit.channel, err = rabbit.conn.Channel()
	rabbit.failOnErr(err, "获取Channel失败")

	return &rabbit
}

// 销毁连接
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
		panic(fmt.Sprintf("%s: %s", message, err))
	}
}


