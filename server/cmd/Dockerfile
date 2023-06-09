FROM golang:latest

WORKDIR /mafia_server

COPY . .
RUN go mod download

RUN cd server/cmd && go install
RUN go build -o mafia_server server/cmd/main.go

CMD ["./mafia_server"]
