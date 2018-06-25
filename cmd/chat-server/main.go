// A simple chat-server over TCP
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// Conn is a generic stream-oriented network connection.
	type Conn net.Conn

	// New connections will be channelled through this one.
	newConnections := make(chan Conn)

	// Send a client to this channel if connection is/should be closed.
	closeConnections := make(chan Conn)

	// Map connections to client IDs.
	clients := make(map[Conn]int)

	// Each client ID is simply an Integer increased by 1.
	clientsIDs := 0

	// Messages will be broadcasested through this channel.
	messages := make(chan string)

	// Port can be set as a param from the command line.
	// Default port is :8080.
	port := flag.String("port", ":8080", "chat server port")
	flag.Parse()

	// ----------------------------------------------------------
	// Start the TCP server
	// ----------------------------------------------------------
	server, err := net.Listen("tcp", *port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Printf("Chat Server is listening on port: %s", *port)
	defer server.Close()

	// ----------------------------------------------------------
	// Accept connections infinitely.
	// ----------------------------------------------------------
	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			// Add a connection to the channel tracking new connections.
			newConnections <- conn
		}
	}()

	// Here's the loop that keeps all connections, messages passing
	// to the right channels.
	for {
		select {
		// Pass messages to connected clients.
		case message := <-messages:

			// To broadcast a message, let's loop through all of them.
			for client := range clients {

				// Send a message.
				go func(conn Conn, msg string) {
					_, err := conn.Write([]byte(msg))

					// Close the connection if an error occured.
					if err != nil {
						closeConnections <- conn
					}

				}(client, message)
			}
			log.Printf("%s[broadcased to %d clients]", message, len(clients))

		case conn := <-newConnections:
			log.Printf("New <Client %d> has joined the chat server.", clientsIDs)

			clients[conn] = clientsIDs
			clientsIDs++

			go func(conn Conn, clientID int) {
				reader := bufio.NewReader(conn)

				for {
					// Read string until the first occurence of '\n'
					newMsg, err := reader.ReadString('\n')
					if err != nil {
						break
					}

					messages <- fmt.Sprintf("<Client %d> %s", clientID, newMsg)
				}

				closeConnections <- conn
			}(conn, clients[conn])

		// Close connections.
		case conn := <-closeConnections:
			log.Printf("<Client %d> disconnected", clients[conn])
			delete(clients, conn)
		}

	}
}
