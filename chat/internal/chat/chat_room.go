package chat

import (
	"fmt"
	"log"
	pb "soa_mafia/chat/proto"

	"github.com/golang/protobuf/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type chatUser struct {
	info      *pb.User
	queueName string
}

type ChatRoom struct {
	Id    string
	users []*chatUser
	ch    *amqp.Channel
}

func NewChatRoom(id string, ch *amqp.Channel) *ChatRoom {
	return &ChatRoom{
		Id:    id,
		users: make([]*chatUser, 0),
		ch:    ch,
	}
}

func (cr *ChatRoom) RegisterUser(user *pb.User) {
	err := cr.ch.ExchangeDeclare(
		cr.Id,    // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := cr.ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = cr.ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		cr.Id,  // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	cr.users = append(cr.users, &chatUser{
		info:      user,
		queueName: q.Name,
	})
}

func (cr *ChatRoom) SendMessage(msg *pb.UserMessage) {
	body, err := proto.Marshal(msg)
	failOnError(err, fmt.Sprintf("Failed to marshal msg: %s", msg.String()))

	err = cr.ch.Publish(
		cr.Id, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}

func (cr *ChatRoom) getChatUser(userId string) *chatUser {
	for _, user := range cr.users {
		if user.info.GetId() == userId {
			return user
		}
	}
	return nil
}

func (cr *ChatRoom) GetMessageChan(userId string) <-chan amqp.Delivery {
	user := cr.getChatUser(userId)
	ch, err := cr.ch.Consume(
		user.queueName, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")
	return ch
}
