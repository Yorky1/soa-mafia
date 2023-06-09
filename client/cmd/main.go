package main

import (
	"log"
	"os"
	"soa_mafia/client/internal/client"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Client running ...")
	conn, err := grpc.Dial(":9001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	_, is_bot := os.LookupEnv("BOT")
	var cli *client.Client
	if is_bot {
		cli = client.NewBotClient(conn)
	} else {
		cli = client.NewRealClient(conn)
	}

	cli.StartGameClient()
}
