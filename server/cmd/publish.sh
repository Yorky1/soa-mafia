#!/bin/sh
docker build -t yorky0/mafia_server -f ./server/cmd/Dockerfile .
docker image push yorky0/mafia_server:latest
