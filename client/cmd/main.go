package main

import (
	"log"
	"os"
	"soa_mafia/client/internal/client"

	"google.golang.org/grpc"
)

func isBot() bool {
	_, isBot := os.LookupEnv("BOT")
	return isBot
}

func getServerPort() string {
	port, has := os.LookupEnv("SERVER_PORT")
	if !has {
		return "9001"
	}
	return port
}

func getServerAddr() string {
	return os.Getenv("SERVER_ADDR")
}

func main() {
	log.Println("Client running ...")
	conn, err := grpc.Dial(getServerAddr()+":"+getServerPort(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	var cli *client.Client
	if isBot() {
		cli = client.NewBotClient(conn)
	} else {
		cli = client.NewRealClient(conn)
	}

	cli.StartGameClient()
}
