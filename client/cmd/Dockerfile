FROM golang:latest

WORKDIR /mafia_client

COPY . .
RUN go mod download

RUN cd client/cmd && go install
RUN go build -o mafia_client client/cmd/main.go

CMD ["./mafia_client"]
