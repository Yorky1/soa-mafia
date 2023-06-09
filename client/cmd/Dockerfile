FROM golang:latest

WORKDIR /mafia_client

COPY . .
RUN go mod download

RUN cd client/cmd && go install
RUN go build -o mafia_client client/cmd/main.go

ENV SERVER_ADDR=mafia_server
ENV SERVER_PORT=9001

CMD ["./mafia_client"]
