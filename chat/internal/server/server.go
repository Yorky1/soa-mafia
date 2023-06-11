package server

import (
	"context"
	"errors"
	"log"

	chat "soa_mafia/chat/internal/chat"
	pb "soa_mafia/chat/proto"

	"github.com/golang/protobuf/proto"
)

type Server struct {
	pb.UnimplementedChatServer
	chatManager *chat.ChatManager
}

func NewServer(host, port, user, pass string) *Server {
	return &Server{
		chatManager: chat.NewChatManager(&chat.AmqpCredentials{
			Host:     host,
			Port:     port,
			User:     user,
			Password: pass,
		}),
	}
}

func (s *Server) RegisterNewRoom(ctx context.Context, request *pb.RegisterRoomRequest) (*pb.Empty, error) {
	log.Printf("RegisterNewRoom request: %s", request.String())
	s.chatManager.RegisterNewRoom(request.RoomId)
	return &pb.Empty{}, nil
}

func (s *Server) EnterRoom(ctx context.Context, request *pb.EnterRoomRequest) (*pb.Empty, error) {
	log.Printf("EnterRoom request: %s", request.String())
	s.chatManager.EnterRoom(request.RoomId, request.User)
	return &pb.Empty{}, nil
}

func (s *Server) Stream(stream pb.Chat_StreamServer) error {
	startAction, err := stream.Recv()
	if err != nil {
		return err
	}
	if startAction.GetType() != pb.UserActionType_Connect {
		return errors.New("No connect message")
	}
	roomId := startAction.GetConnect().GetRoomId()
	userId := startAction.GetConnect().GetUserId()

	log.Printf("Stream request: %s", startAction.String())
	go func() error {
		for {
			action, err := stream.Recv()
			if err != nil {
				return err
			}
			switch action.GetType() {
			case pb.UserActionType_Message:
				log.Printf("Message received: %s", action.GetMessage())
				s.chatManager.SendMessage(roomId, action.GetMessage())
			case pb.UserActionType_Connect:
				log.Printf("Connect action received: %s", action.GetConnect())
				return errors.New("already connected")
			}
		}
	}()

	ch := s.chatManager.GetMessageChan(roomId, userId)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-ch:
			pbMsg := &pb.UserMessage{}
			proto.Unmarshal(msg.Body, pbMsg)
			log.Printf("Send message to %s: %s", userId, pbMsg.String())
			stream.Send(&pb.ServerChatAction{
				Action: &pb.ServerChatAction_Message{
					Message: pbMsg,
				},
			})
		}
	}
}
