
FROM golang:1.10
LABEL maintainer="web@rafalgolarz.com"

COPY cmd/chat-server/  /go/src/github.com/rafalgolarz/tcpchat/cmd/chat-server/
COPY vendor/ /go/src/
WORKDIR /go/src/github.com/rafalgolarz/tcpchat/cmd/chat-server
RUN go build
RUN chmod +x chat-server

CMD ./chat-server
EXPOSE 8080