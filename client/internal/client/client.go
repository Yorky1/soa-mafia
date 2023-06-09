package client

import (
	"context"
	"fmt"
	"log"
	pb "soa_mafia/server/proto"

	"google.golang.org/grpc"
)

type Client struct {
	player    *pb.Player
	sessionId string
	pb        pb.GameClient
	voter     Voter
}

func NewRealClient(conn *grpc.ClientConn) *Client {
	return &Client{pb: pb.NewGameClient(conn), voter: NewRealVoter()}
}

func NewBotClient(conn *grpc.ClientConn) *Client {
	return &Client{pb: pb.NewGameClient(conn), voter: NewBotVoter()}
}

func (client *Client) register() {
	fmt.Println(HELLO_MESSAGE)
	clientName := client.voter.GetName()

	request := &pb.RegisterRequest{PlayerName: clientName}
	response, err := client.pb.Register(context.TODO(), request)
	if err != nil {
		log.Fatalf("client.Register failed: %v", err)
		panic(err)
	}
	client.sessionId = response.GameSessionId
	client.player = response.CurrentPlayer
}

func (client *Client) gameInfo() *pb.GameInfo {
	request := &pb.GameSession{
		Id: client.sessionId,
	}
	gameInfo, err := client.pb.CurGameInfo(context.TODO(), request)
	if err != nil {
		log.Fatalf("client.Register failed: %v", err)
		panic(err)
	}
	return gameInfo
}

func (client *Client) waitForNextState() {
	stream, err := client.pb.WaitForGame(context.TODO(), &pb.GameSession{Id: client.sessionId})
	if err != nil {
		log.Fatalf("client.WaitForGame failed: %v", err)
		panic(err)
	}
	stream.Recv()
}

func (client *Client) waitForGameStart() {
	fmt.Println("Cool! Now we are waiting for the other players to connect...")
	client.waitForNextState()

	fmt.Println("Ok, all players are connected. Let's start!")
	fmt.Println("Connected players are:")
	players := client.gameInfo().Players
	for _, player := range players {
		fmt.Println(player.Name)
	}
	client.updateRole(players)
}

func getPlayersNamesWithIdToVote(players []*pb.Player) []string {
	res := make([]string, 0)
	for i, player := range players {
		if player.Role != "ghost" {
			res = append(res, fmt.Sprintf("%d %s", i, player.Name))
		}
	}
	return res
}

func findPlayer(players []*pb.Player, vote string) *pb.Player {
	for i, player := range players {
		if fmt.Sprintf("%d %s", i, player.Name) == vote {
			return player
		}
	}
	return nil
}

func (client *Client) updateRole(players []*pb.Player) {
	for _, player := range players {
		if player.Id == client.player.Id {
			client.player = player
			return
		}
	}
}

func (client *Client) waitForDayVote() bool {
	gameInfo := client.gameInfo()
	fmt.Printf("Current day is %d\n", gameInfo.Day)
	fmt.Printf("Your role is %s\n", client.player.Role)
	if client.player.Role != "ghost" {
		vote := client.voter.DayVote(getPlayersNamesWithIdToVote(gameInfo.Players))

		if _, err := client.pb.VotePlayer(context.TODO(), &pb.PlayerVote{
			PlayerId:  findPlayer(gameInfo.Players, vote).Id,
			SessionId: client.sessionId,
		}); err != nil {
			log.Fatalf("client.VotePlayer failed: %v", err)
			panic(err)
		}
	} else {
		fmt.Println("Sorry, but you dead, so can't vote :(")
	}

	client.waitForNextState()

	result, err := client.pb.StageResult(context.TODO(), &pb.GameSession{Id: client.sessionId})
	if err != nil {
		log.Fatalf("client.StageResult failed: %v", err)
		panic(err)
	}
	fmt.Println(result.Message)
	return result.GameOver
}

func (client *Client) waitForSheriff() bool {
	if client.player.Role == "sheriff" {
		gameInfo := client.gameInfo()
		vote := client.voter.SheriffVote(getPlayersNamesWithIdToVote(gameInfo.Players))

		if _, err := client.pb.VotePlayer(context.TODO(), &pb.PlayerVote{
			PlayerId:  findPlayer(gameInfo.Players, vote).Id,
			SessionId: client.sessionId,
		}); err != nil {
			log.Fatalf("client.VotePlayer failed: %v", err)
			panic(err)
		}
	}
	client.waitForNextState()
	fmt.Println("Sheriff made his choise. Mafia wakes up!")
	return false
}

func (client *Client) waitForMafia() bool {
	if client.player.Role == "mafia" {
		gameInfo := client.gameInfo()
		vote := client.voter.MafiaVote(getPlayersNamesWithIdToVote(gameInfo.Players))

		if _, err := client.pb.VotePlayer(context.TODO(), &pb.PlayerVote{
			PlayerId:  findPlayer(gameInfo.Players, vote).Id,
			SessionId: client.sessionId,
		}); err != nil {
			log.Fatalf("client.VotePlayer failed: %v", err)
			panic(err)
		}
	}
	client.waitForNextState()
	result, err := client.pb.StageResult(context.TODO(), &pb.GameSession{Id: client.sessionId})
	if err != nil {
		log.Fatalf("client.StageResult failed: %v", err)
		panic(err)
	}
	fmt.Println(result.Message)
	client.updateRole(result.Players)
	return result.GameOver
}

func (client *Client) StartGameClient() {
	client.register()
	client.waitForGameStart()

	for {
		if client.waitForDayVote() {
			break
		}
		if client.waitForSheriff() {
			break
		}
		if client.waitForMafia() {
			break
		}
	}
}

const HELLO_MESSAGE = `Greetings! You are playing a SOA-Mafia game.`
