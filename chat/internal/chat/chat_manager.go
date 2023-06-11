package chat

import (
	"fmt"
	"log"
	pb "soa_mafia/chat/proto"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpCredentials struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (ac *AmqpCredentials) String() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", ac.User, ac.Password, ac.Host, ac.Port)
}

type ChatManager struct {
	rooms map[string]*ChatRoom
	conn  *amqp.Connection
	ch    *amqp.Channel
}

func NewChatManager(cred *AmqpCredentials) *ChatManager {
	log.Printf("Trying to connect to %s", cred.String())
	conn, err := amqp.Dial(cred.String())
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	log.Println("Channel created")
	return &ChatManager{
		rooms: make(map[string]*ChatRoom),
		conn:  conn,
		ch:    ch,
	}
}

func (cm *ChatManager) RegisterNewRoom(roomId string) {
	if _, has := cm.rooms[roomId]; !has {
		cm.rooms[roomId] = NewChatRoom(roomId, cm.ch)
	}
}

func (cm *ChatManager) EnterRoom(roomId string, user *pb.User) {
	cm.rooms[roomId].RegisterUser(user)
}

func (cm *ChatManager) SendMessage(roomId string, userMessage *pb.UserMessage) {
	cm.rooms[roomId].SendMessage(userMessage)
}

func (cm *ChatManager) GetMessageChan(roomId string, userId string) <-chan amqp.Delivery {
	return cm.rooms[roomId].GetMessageChan(userId)
}
