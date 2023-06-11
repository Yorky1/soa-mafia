#!/bin/sh
docker build -t yorky0/chat_server -f ./chat/cmd/Dockerfile .
docker image push yorky0/chat_server:latest
