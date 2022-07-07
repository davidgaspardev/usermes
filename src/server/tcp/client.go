package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	log "github.com/davidgaspardev/golog"
)

// Client
type _Client struct {
	conn    net.Conn
	command chan<- _Command
}

// Send data to client
func (client *_Client) send(data []byte) {
	if data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	client.conn.Write(data)
}

// Listen data from client
func (client *_Client) listen() error {
	for {
		data, err := bufio.NewReader(client.conn).ReadBytes('\n')
		if err != nil {
			log.Error("TCP Client", fmt.Errorf("connection missed: %s", client.conn.RemoteAddr().String()))
			return err
		}

		// Load data raw from client
		dataRaw := string(data[:len(data)-1])

		args := strings.Split(dataRaw, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/record":
			client.command <- _Command{
				id:     CMD_RECORD,
				client: client,
				args:   args,
			}
		default:
			client.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

// Send error to client
func (client *_Client) err(err error) {
	log.Error("TCP Client", err)
	client.send([]byte(err.Error()))
}
