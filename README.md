## TCP CHAT

An example of a simple TCP chat server and client.

## Running the server

Run the chat server on the default port :8080

```!/bin/bash
go run cmd/chat-server/main.go
```

Run the chat server on the custom port (for instance :9090)

```!/bin/bash
go run cmd/chat-server/main.go -port=:9090
```

You can also run it using Docker:

```!/bin/bash
docker build -f Dockerfile -t rafalgolarz/tcpchat .
docker run --rm -p 8080:8080 --name=tcpchat rafalgolarz/tcpchat
```

## Running the client

Run the chat client on the default port :8080

```!/bin/bash
go run cmd/chat-client/main.go
```

Run the chat client and connect it to a custom port (for instance :9090)

```!/bin/bash
go run cmd/chat-client/main.go -port=:9090
```

Run the chat client and connect it to a custom host/ip and port (for instance 192.168.1.11:9090)

```!/bin/bash
go run cmd/chat-client/main.go -host=192.168.1.11 -port=:9090
```