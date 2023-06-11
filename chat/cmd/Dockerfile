FROM golang:latest

WORKDIR /chat_server

COPY . .
RUN go mod download

RUN cd chat/cmd && go install
RUN go build -o chat_server chat/cmd/main.go

CMD ["./chat_server"]
