package gamesession

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"sync/atomic"

	"github.com/google/uuid"
)

type GameState int32

const (
	Day GameState = iota
	SheriffChooses
	MafiaChooses
)

type GameSessionSnapshot struct {
	wg           sync.WaitGroup
	currentVotes []string
}

func NewGameSessionSnapshot() *GameSessionSnapshot {
	return &GameSessionSnapshot{
		currentVotes: make([]string, 0),
	}
}

type GameSession struct {
	id        string
	players   []*Player
	snapshots []*GameSessionSnapshot
	day       int
	gameState GameState
	state     int32
	stageInfo string
}

func NewGameSession(maxPlayers int) *GameSession {
	gs := &GameSession{
		id:        uuid.New().String(),
		players:   make([]*Player, 0),
		snapshots: make([]*GameSessionSnapshot, 0),
		day:       -1,
		gameState: Day,
		state:     0,
		stageInfo: "",
	}
	gs.snapshots = append(gs.snapshots, NewGameSessionSnapshot())
	gs.snapshots = append(gs.snapshots, NewGameSessionSnapshot())
	gs.snapshots[0].wg.Add(maxPlayers)
	return gs
}

func (gs *GameSession) PlayersCount() int {
	return len(gs.players)
}

func (gs *GameSession) Id() string {
	return gs.id
}

func (gs *GameSession) AddPlayer(name string) string {
	gs.players = append(gs.players, newPlayer(name))
	return gs.players[len(gs.players)-1].Id
}

func (gs *GameSession) roleCount(role PlayerRole) int {
	res := 0
	for _, player := range gs.players {
		if player.Role == role {
			res++
		}
	}
	return res
}

func (gs *GameSession) nonGhostCount() int {
	return gs.roleCount(Citizen) + gs.roleCount(Mafia) + gs.roleCount(Sheriff)
}

func (gs *GameSession) initGame() {
	playerCount := len(gs.players)
	mafiaCount := playerCount / 3
	sheriffCount := 0
	if playerCount > 3 {
		sheriffCount++
	}

	ids := make([]int, 0)
	for i := range gs.players {
		ids = append(ids, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
	for i := range ids {
		player := gs.players[ids[i]]
		if mafiaCount > 0 {
			player.Role = Mafia
			mafiaCount--
		} else if sheriffCount > 0 {
			player.Role = Sheriff
			sheriffCount--
		} else {
			player.Role = Citizen
		}
	}

	log.Println("Init game, roles selected")
	for _, player := range gs.players {
		log.Printf("%s %s", player.Name, player.Role)
	}

	gs.day = 0
	gs.gameState = Day
	gs.snapshots = append(gs.snapshots, NewGameSessionSnapshot())
	gs.snapshots[gs.state].wg.Add(len(gs.players))
	log.Printf("Init game state: %d\n", gs.state)
}

func (gs *GameSession) GameOver() bool {
	return gs.roleCount(Mafia) >= gs.roleCount(Citizen)+gs.roleCount(Sheriff) || gs.roleCount(Mafia) == 0
}

func (gs *GameSession) GameOverMessage() string {
	res := "Suddenly! The game is over.\n"
	if gs.roleCount(Mafia) > 0 {
		res += "Mafia wins!"
	} else {
		res += "Citizen win!"
	}
	return res
}

func (gs *GameSession) VotePlayer(playerId string) {
	sh := gs.snapshots[gs.state]
	sh.currentVotes = append(sh.currentVotes, playerId)
}

func (gs *GameSession) findById(playerId string) *Player {
	for _, player := range gs.players {
		if player.Id == playerId {
			return player
		}
	}
	panic("findById failed!")
}

func (gs *GameSession) StageResultMessage() string {
	return gs.stageInfo
}

func (gs *GameSession) makeDayVoting() {
	sh := gs.snapshots[gs.state-1]
	mp := make(map[string]int)
	for _, vote := range sh.currentVotes {
		mp[vote]++
	}

	cnt := gs.nonGhostCount()
	victim := ""
	for key, value := range mp {
		if cnt < value*2 {
			victim = key
		}
	}

	if victim != "" {
		player := gs.findById(victim)
		gs.stageInfo = fmt.Sprintf("According to the results of the vote, it was decided to execute the player %s, he was %s.", player.Name, player.Role)
		player.Role = Ghost
	} else {
		gs.stageInfo = "People didn't agree on which player to execute today :("
	}

	if gs.GameOver() {
		gs.stageInfo += "\n"
		gs.stageInfo += gs.GameOverMessage()
	} else {
		gs.stageInfo += "\nNight is coming! Sheriff wakes up."
	}

	gs.gameState = SheriffChooses
	gs.snapshots = append(gs.snapshots, NewGameSessionSnapshot())
	gs.snapshots[gs.state].wg.Add(len(gs.players))
}

func (gs *GameSession) makeSheriffVoting() {
	sh := gs.snapshots[gs.state-1]

	if gs.roleCount(Sheriff) > 0 {
		vote := sh.currentVotes[0]
		player := gs.findById(vote)
		gs.stageInfo = fmt.Sprintf("Tonight the sheriff checked %s, he turned out to be %s.\n", player.Name, player.Role)
	} else {
		gs.stageInfo = fmt.Sprintf("The sheriff is either dead or decided not to check on anyone tonight.\n")
	}

	gs.gameState = MafiaChooses
	gs.snapshots = append(gs.snapshots, NewGameSessionSnapshot())
	gs.snapshots[gs.state].wg.Add(len(gs.players))
}

func (gs *GameSession) makeMafiaVoting() {
	log.Println("makeMafiaVoting started")
	sh := gs.snapshots[gs.state-1]

	votes := sh.currentVotes
	all_equal := true
	for i := range votes {
		if i > 0 && votes[i] != votes[i-1] {
			all_equal = false
		}
	}
	if len(votes) == 0 {
		all_equal = false
	}

	if all_equal {
		victim := votes[0]
		player := gs.findById(victim)
		gs.stageInfo += fmt.Sprintf("Mafia killed %s, he was %s.", player.Name, player.Role)
		player.Role = Ghost
	} else {
		gs.stageInfo += "Mafia hasn't come to an agreement on who to kill."
	}

	if gs.GameOver() {
		gs.stageInfo += "\n"
		gs.stageInfo += gs.GameOverMessage()
	}

	gs.gameState = Day
	gs.snapshots = append(gs.snapshots, NewGameSessionSnapshot())
	gs.snapshots[gs.state].wg.Add(len(gs.players))
	gs.day += 1
}

func (gs *GameSession) updateState(oldState int32) {
	if !atomic.CompareAndSwapInt32(&gs.state, oldState, oldState+1) {
		return
	}

	if gs.day == -1 {
		gs.initGame()
		return
	}

	if gs.gameState == Day {
		gs.makeDayVoting()
		return
	}

	if gs.gameState == SheriffChooses {
		gs.makeSheriffVoting()
		return
	}

	gs.makeMafiaVoting()
}

func (gs *GameSession) NextState() {
	curState := atomic.LoadInt32(&gs.state)
	gs.snapshots[curState].wg.Done()
	gs.snapshots[curState].wg.Wait()
	gs.updateState(curState)
}

func (gs *GameSession) GameInfo() ([]*Player, int) {
	return gs.players, gs.day
}
