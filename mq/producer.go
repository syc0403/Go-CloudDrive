package mq

import (
	"Go-CloudDrive/config"
	"github.com/streadway/amqp"
	"log"
)

/**
* @Description
* @Author YuCheng
* @Date 2023/1/18 22:58
**/
var conn *amqp.Connection
var channel *amqp.Channel

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

func init() {
	// 是否开启异步转移功能，开启时才初始化rabbitMQ连接
	if !config.AsyncTransferEnable {
		return
	}
	if initChannel() {
		channel.NotifyClose(notifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel()
			}
		}
	}()
}

func initChannel() bool {
	//1.判断channel是否已经创建过
	if channel != nil {
		return true
	}
	//2.获得rabbitmq的一个连接
	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//3.打开一个channel，用于消息的发布与接受
	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// Publish : 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel() {
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条小心
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}
