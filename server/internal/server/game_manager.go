package server

import (
	"errors"
	gs "soa_mafia/server/internal/game_session"

	"log"
)

type GameManager struct {
	sessions   []*gs.GameSession
	maxPlayers int
}

func NewGameManager(maxPlayers int) *GameManager {
	return &GameManager{
		sessions: make([]*gs.GameSession, 0),
	}
}

func (gm *GameManager) findSessionToConnect() *gs.GameSession {
	for _, s := range gm.sessions {
		if s.PlayersCount() < gm.maxPlayers {
			return s
		}
	}
	gm.sessions = append(gm.sessions, gs.NewGameSession())
	return gm.sessions[len(gm.sessions)-1]
}

func (gm *GameManager) RegisterPlayer(name string) (string, string) {
	session := gm.findSessionToConnect()
	playerId := session.AddPlayer(name)
	return session.Id(), playerId
}

func (gm *GameManager) findSession(id string) *gs.GameSession {
	for _, s := range gm.sessions {
		if s.Id() == id {
			return s
		}
	}
	return nil
}

func (gm *GameManager) WaitForGame(sessionId string) error {
	s := gm.findSession(sessionId)
	if s == nil {
		log.Fatalf("Session with id %s not found", sessionId)
		return errors.New("session not found")
	}
	s.NextState()
	return nil
}

func (gm *GameManager) GameInfo(sessionId string) ([]*gs.Player, int, error) {
	s := gm.findSession(sessionId)
	if s == nil {
		log.Fatalf("Session with id %s not found", sessionId)
		return nil, 0, errors.New("session not found")
	}
	players, day := s.GameInfo()
	return players, day, nil
}

func (gm *GameManager) StageResult(sessionId string) (string, []*gs.Player, bool) {
	s := gm.findSession(sessionId)
	if s == nil {
		log.Fatalf("Session with id %s not found", sessionId)
	}
	players, _ := s.GameInfo()
	return s.StageResultMessage(), players, s.GameOver()
}

func (gm *GameManager) VotePlayer(sessionId string, playerId string) error {
	s := gm.findSession(sessionId)
	if s == nil {
		log.Fatalf("Session with id %s not found", sessionId)
		return errors.New("session not found")
	}
	s.VotePlayer(playerId)
	return nil
}
