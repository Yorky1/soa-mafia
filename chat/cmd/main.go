package main

import (
	"log"
	"net"
	"os"
	"soa_mafia/chat/internal/server"
	pb "soa_mafia/chat/proto"

	"google.golang.org/grpc"
)

func getEnvOrDefault(env, def string) string {
	res, has := os.LookupEnv(env)
	if has {
		return res
	} else {
		return def
	}
}

func main() {
	lis, err := net.Listen("tcp", ":"+getEnvOrDefault("CHAT_SERVER_PORT", "9002"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()

	pb.RegisterChatServer(srv, server.NewServer(
		getEnvOrDefault("RABBITMQ_HOST", "localhost"),
		getEnvOrDefault("REBBITMQ_PORT", "5672"),
		getEnvOrDefault("RABBITMQ_USER", "guest"),
		getEnvOrDefault("RABBITMQ_PASSWORD", "guest"),
	))
	log.Fatalln(srv.Serve(lis))
}
