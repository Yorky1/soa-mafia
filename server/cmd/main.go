package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"soa_mafia/server/internal/server"

	pb "soa_mafia/server/proto"
)

func getMaxPlayers() int {
	maxPlayers, has := os.LookupEnv("MAX_PLAYERS")
	if !has {
		return 4
	}
	res, err := strconv.Atoi(maxPlayers)
	if err != nil {
		panic(fmt.Sprintf("Bad env_var MAX_PLAYERS: %s %s", maxPlayers, err))
	}
	return res
}

func getServerPort() string {
	port, has := os.LookupEnv("SERVER_PORT")
	if !has {
		return "9001"
	}
	return port
}

func main() {
	lis, err := net.Listen("tcp", ":"+getServerPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()

	pb.RegisterGameServer(srv, server.NewServer(getMaxPlayers()))
	log.Fatalln(srv.Serve(lis))
}
