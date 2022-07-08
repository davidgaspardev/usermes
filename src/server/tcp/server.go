package tcp

import (
	"bytes"
	"fmt"
	"net"

	log "github.com/davidgaspardev/golog"
)

type _Server struct {
	clients  map[[40]byte]*_Client
	commands chan _Command
	orders   chan []byte
}

var server *_Server

func getServer() *_Server {
	if server == nil {
		server = &_Server{
			clients:  make(map[[40]byte]*_Client),
			commands: make(chan _Command),
			orders:   make(chan []byte),
		}
	}

	return server
}

func SendCommand(command []byte) {
	server.orders <- command
}

func (server *_Server) run() {
	server.listenExternalCommand()

	for {
		command := <-server.commands
		log.Info("TCP Server", fmt.Sprint(command))

		switch command.id {
		case CMD_REGISTER:
			// Register raspberry
			server.registerClient(command.client)
		case CMD_RECORD:
			// Record signal
		}
	}
}

func (server *_Server) listenExternalCommand() {
	go func() {
		for {
			orders := <-server.orders

			if bytes.Compare(orders, []byte("start-setup")) == 0 || bytes.Compare(orders, []byte("end-setup")) == 0 {
				for _, client := range server.clients {
					client.send(orders)
				}
			}
		}
	}()
}

// Connect client to network
func (server *_Server) connectClient(conn net.Conn) {
	client := &_Client{
		conn:    conn,
		command: server.commands,
	}

	if err := client.listen(); err != nil {
		server.unregisterClient(client)
	}
}

func (server *_Server) registerClient(client *_Client) {
	server.clients[client.hash] = client
}

func (server *_Server) unregisterClient(client *_Client) {
	delete(server.clients, client.hash)
}
