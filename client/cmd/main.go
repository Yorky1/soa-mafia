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

func getEnvOrDefault(env, def string) string {
	res, has := os.LookupEnv(env)
	if !has {
		return def
	} else {
		return res
	}
}

func main() {
	log.Println("Client running ...")
	conn, err := grpc.Dial(getEnvOrDefault("SERVER_ADDR", "localhost")+":"+getEnvOrDefault("SERVER_PORT", "9001"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	chat_conn, err := grpc.Dial(getEnvOrDefault("CHAT_ADDR", "localhost")+":"+getEnvOrDefault("CHAT_PORT", "9002"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer chat_conn.Close()

	var cli *client.Client
	if isBot() {
		cli = client.NewBotClient(conn, chat_conn)
	} else {
		cli = client.NewRealClient(conn, chat_conn)
	}

	cli.StartGameClient()
}
