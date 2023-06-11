#!/bin/sh
docker build -t yorky0/mafia_client -f ./client/cmd/Dockerfile .
docker image push yorky0/mafia_client:latest
