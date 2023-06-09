package server

import (
	"context"
	"fmt"
	"log"

	gs "soa_mafia/server/internal/game_session"
	pb "soa_mafia/server/proto"
)

type Server struct {
	pb.UnimplementedGameServer
	gameManager *GameManager
}

func NewServer(maxPlayers int) *Server {
	return &Server{
		gameManager: NewGameManager(maxPlayers),
	}
}

func (s *Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	name := request.GetPlayerName()
	log.Println(fmt.Sprintf("Register request: %s", name))
	sessionId, playerId := s.gameManager.RegisterPlayer(name)
	log.Println(fmt.Sprintf("GameSessionId: %s", sessionId))
	return &pb.RegisterResponse{
		GameSessionId: sessionId,
		CurrentPlayer: &pb.Player{
			Name: name,
			Id:   playerId,
			Role: gs.Ghost.String(),
		},
	}, nil
}

func (s *Server) WaitForGame(gs *pb.GameSession, stream pb.Game_WaitForGameServer) error {
	log.Println(fmt.Sprintf("WaitForGame request: %s", gs.GetId()))
	err := s.gameManager.WaitForGame(gs.GetId())
	if err != nil {
		return err
	}
	err = stream.Send(&pb.Empty{})
	return err
}

func (s *Server) CurGameInfo(ctx context.Context, request *pb.GameSession) (*pb.GameInfo, error) {
	log.Println(fmt.Sprintf("CurGameInfo request: %s", request.GetId()))
	players, day, err := s.gameManager.GameInfo(request.GetId())
	if err != nil {
		return nil, err
	}
	pbPlayers := make([]*pb.Player, 0)
	for _, player := range players {
		pbPlayers = append(pbPlayers, &pb.Player{Name: player.Name, Id: player.Id, Role: player.Role.String()})
	}
	return &pb.GameInfo{
		Players: pbPlayers,
		Day:     int32(day),
	}, nil
}

func (s *Server) VotePlayer(ctx context.Context, vote *pb.PlayerVote) (*pb.Empty, error) {
	log.Println(fmt.Sprintf("VotePlayer request: %s", vote.GetSessionId()))
	err := s.gameManager.VotePlayer(vote.SessionId, vote.PlayerId)
	return &pb.Empty{}, err
}

func (s *Server) StageResult(ctx context.Context, gs *pb.GameSession) (*pb.VotingResult, error) {
	log.Println(fmt.Sprintf("StageResult request: %s", gs.GetId()))
	msg, players, gameOver := s.gameManager.StageResult(gs.GetId())
	pbPlayers := make([]*pb.Player, 0)
	for _, player := range players {
		pbPlayers = append(pbPlayers, &pb.Player{Name: player.Name, Id: player.Id, Role: player.Role.String()})
	}
	result := &pb.VotingResult{
		Message:  msg,
		Players:  pbPlayers,
		GameOver: gameOver,
	}
	log.Printf("StageResult: %s\n", result)
	return result, nil
}
