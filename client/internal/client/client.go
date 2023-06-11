package client

import (
	"context"
	"fmt"
	"log"
	chatpb "soa_mafia/chat/proto"
	pb "soa_mafia/server/proto"

	"github.com/pterm/pterm"
	"google.golang.org/grpc"
)

type Client struct {
	player      *pb.Player
	sessionId   string
	pb          pb.GameClient
	chat        chatpb.ChatClient
	voter       Voter
	chatBuffers map[string][]*chatpb.UserMessage
}

func NewRealClient(conn *grpc.ClientConn, chat_conn *grpc.ClientConn) *Client {
	return &Client{pb: pb.NewGameClient(conn), chat: chatpb.NewChatClient(chat_conn), voter: NewRealVoter(), chatBuffers: make(map[string][]*chatpb.UserMessage)}
}

func NewBotClient(conn *grpc.ClientConn, chat_conn *grpc.ClientConn) *Client {
	return &Client{pb: pb.NewGameClient(conn), chat: chatpb.NewChatClient(chat_conn), voter: NewBotVoter(), chatBuffers: make(map[string][]*chatpb.UserMessage)}
}

func (client *Client) register() {
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
	pterm.Println("Now we are waiting for the other players to connect...")
	client.waitForNextState()

	pterm.Println("Ok, all players are connected. Let's start!")
	pterm.Println("Connected players are:")
	players := client.gameInfo().Players
	for _, player := range players {
		pterm.Println(player.Name)
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

func (client *Client) votingCycle(chatName string, startMsg string, voteFunc func()) {
	for {
		choise := client.voter.GetVoteMenuChoise()
		switch choise {
		case GoToChat:
			client.chatEngine(chatName, startMsg)
		case Vote:
			voteFunc()
			return
		}
	}
}

func (client *Client) waitForDayVote() bool {
	gameInfo := client.gameInfo()
	pterm.Printf("Current day is %d\n", gameInfo.Day)
	pterm.Printf("Your role is %s\n", client.player.Role)
	if client.player.Role != "ghost" {
		client.votingCycle(
			client.sessionId+"-day",
			"Welcome to day chat. Here you can chat with all players with non-ghost role.",
			func() {
				vote := client.voter.DayVote(getPlayersNamesWithIdToVote(gameInfo.Players))

				if _, err := client.pb.VotePlayer(context.TODO(), &pb.PlayerVote{
					PlayerId:  findPlayer(gameInfo.Players, vote).Id,
					SessionId: client.sessionId,
				}); err != nil {
					log.Fatalf("client.VotePlayer failed: %v", err)
					panic(err)
				}
			},
		)
	} else {
		pterm.Println("Sorry, but you dead, so can't vote or chat with people :(")
	}

	client.waitForNextState()

	result, err := client.pb.StageResult(context.TODO(), &pb.GameSession{Id: client.sessionId})
	if err != nil {
		log.Fatalf("client.StageResult failed: %v", err)
		panic(err)
	}
	pterm.Println(result.Message)
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
	pterm.Println("Sheriff made his choise. Mafia wakes up!")
	return false
}

func (client *Client) waitForMafia() bool {
	if client.player.Role == "mafia" {
		client.votingCycle(
			client.sessionId+"-mafia",
			"Welcome to mafia chat! Here you can chat with other players with role mafia.",
			func() {
				gameInfo := client.gameInfo()
				vote := client.voter.MafiaVote(getPlayersNamesWithIdToVote(gameInfo.Players))

				if _, err := client.pb.VotePlayer(context.TODO(), &pb.PlayerVote{
					PlayerId:  findPlayer(gameInfo.Players, vote).Id,
					SessionId: client.sessionId,
				}); err != nil {
					log.Fatalf("client.VotePlayer failed: %v", err)
					panic(err)
				}
			},
		)
	}
	client.waitForNextState()
	result, err := client.pb.StageResult(context.TODO(), &pb.GameSession{Id: client.sessionId})
	if err != nil {
		log.Fatalf("client.StageResult failed: %v", err)
		panic(err)
	}
	pterm.Println(result.Message)
	client.updateRole(result.Players)
	return result.GameOver
}

func (client *Client) getChatBuffer(chatName string) []*chatpb.UserMessage {
	if _, has := client.chatBuffers[chatName]; !has {
		client.chatBuffers[chatName] = make([]*chatpb.UserMessage, 0)
	}
	return client.chatBuffers[chatName]
}

func formatChatMessage(name string, msg string) string {
	return fmt.Sprintf("%s: %s", name, msg)
}

func (client *Client) getChatFrame(chatBuffer []*chatpb.UserMessage) string {
	res := ""
	for i := len(chatBuffer) - 5; i < len(chatBuffer); i++ {
		if i >= 0 {
			res += pterm.Sprintln("%s: %s", chatBuffer[i].User, chatBuffer[i].Message)
		}
	}
	return res
}

func (client *Client) chatEngine(chatName string, startMsg string) {
	pterm.Println(startMsg)
	pterm.Println(`Type /exit to leave chat room`)

	client.chat.RegisterNewRoom(context.TODO(), &chatpb.RegisterRoomRequest{
		RoomId: chatName,
	})

	client.chat.EnterRoom(context.TODO(), &chatpb.EnterRoomRequest{
		RoomId: chatName,
		User: &chatpb.User{
			Id:   client.player.Id,
			Name: client.player.Name,
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.chat.Stream(ctx)
	if err != nil {
		log.Panic(err)
	}

	stream.Send(&chatpb.UserChatAction{
		Action: &chatpb.UserChatAction_Connect{
			Connect: &chatpb.ConnectionInfo{
				UserId: client.player.Id,
				RoomId: chatName,
			},
		},
		Type: chatpb.UserActionType_Connect,
	})

	chatBuffer := client.getChatBuffer(chatName)
	pterm.Println(client.getChatFrame(chatBuffer))

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				return
			}
			userMessage := msg.GetMessage()
			chatBuffer = append(chatBuffer, userMessage)
			pterm.Println(formatChatMessage(msg.GetMessage().User.Name, msg.GetMessage().Message))
		}
	}()

	for {
		msg := client.voter.GetMessage()
		if msg == "/exit" {
			break
		}
		stream.Send(&chatpb.UserChatAction{
			Action: &chatpb.UserChatAction_Message{
				Message: &chatpb.UserMessage{
					Message: msg,
					User: &chatpb.User{
						Id:   client.player.Id,
						Name: client.player.Name,
					},
				},
			},
			Type: chatpb.UserActionType_Message,
		})
	}
	stream.CloseSend()
}

func (client *Client) mainGameCycle() {
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

func (client *Client) showMainMenu() {
	shutDown := false
	for !shutDown {
		switch client.voter.GetMainMenuChoise() {
		case FindGameSession:
			client.mainGameCycle()
		case Exit:
			shutDown = true
			break
		}
	}
}

func (client *Client) StartGameClient() {
	pterm.Print(GAME_LOGO)
	pterm.Println(HELLO_MESSAGE)

	client.showMainMenu()
}

const HELLO_MESSAGE = `Greetings! You are playing a SOA-Mafia game.`

const GAME_LOGO = `
███████╗ ██████╗  █████╗       ███╗   ███╗ █████╗ ███████╗██╗ █████╗ 
██╔════╝██╔═══██╗██╔══██╗      ████╗ ████║██╔══██╗██╔════╝██║██╔══██╗
███████╗██║   ██║███████║█████╗██╔████╔██║███████║█████╗  ██║███████║
╚════██║██║   ██║██╔══██║╚════╝██║╚██╔╝██║██╔══██║██╔══╝  ██║██╔══██║
███████║╚██████╔╝██║  ██║      ██║ ╚═╝ ██║██║  ██║██║     ██║██║  ██║
╚══════╝ ╚═════╝ ╚═╝  ╚═╝      ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝                                                                     
`
