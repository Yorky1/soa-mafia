package gamesession

import "github.com/google/uuid"

type PlayerRole int64

const (
	Citizen PlayerRole = iota
	Mafia
	Sheriff
	Ghost
)

func (pr PlayerRole) String() string {
	switch pr {
	case Citizen:
		return "citizen"
	case Mafia:
		return "mafia"
	case Sheriff:
		return "sheriff"
	case Ghost:
		return "ghost"
	}
	return "unknown"
}

type Player struct {
	Name string
	Id   string
	Role PlayerRole
}

func newPlayer(name string) *Player {
	return &Player{
		Name: name,
		Id:   uuid.New().String(),
		Role: Ghost,
	}
}
