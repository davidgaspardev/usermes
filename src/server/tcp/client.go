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
	hash    [40]byte
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
func (client *_Client) listen() (err error) {
	firstCommand := true

	for {
		var data []byte
		data, err = bufio.NewReader(client.conn).ReadBytes('\n')
		if err != nil {
			log.Error("TCP Client", fmt.Errorf("connection missed: %s", client.conn.RemoteAddr().String()))
			return err
		}

		// Load data raw from client
		// dataRaw := string(data[:len(data)-1])

		// args := strings.Split(dataRaw, " ")
		// cmd := strings.TrimSpace(args[0])
		var cmd string
		var args []string
		cmd, args, err = client.loadCommand(data)

		if err != nil {
			client.quit()
			return
		}

		if firstCommand {
			firstCommand = false

			if cmd != "/register" {
				client.quit()
				return
			}

			client.setHash(args[0])

			// Send command to server
			client.command <- _Command{
				id:     CMD_REGISTER,
				client: client,
				args:   args,
			}

			continue
		}

		switch cmd {
		case "/record":
			// Send command to server
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

func (client *_Client) loadCommand(data []byte) (cmd string, args []string, err error) {
	dataRaw := string(data[:len(data)-1])
	dataPart := strings.Split(dataRaw, " ")

	if len(dataPart) < 2 {
		err = fmt.Errorf("data invalid")
		return
	}

	cmd = strings.TrimSpace(dataPart[0])
	args = make([]string, len(dataPart)-1)

	for i := 0; i < len(args); i++ {
		args[i] = dataPart[i+1]
	}

	return
}

// Send error to client
func (client *_Client) err(err error) {
	log.Error("TCP Client", err)
	client.send([]byte(err.Error()))
}

func (client *_Client) quit() {
	log.System("TCP Client", "connection finished: "+client.conn.RemoteAddr().String())
	client.conn.Close()
}

func (client *_Client) setHash(hash string) error {
	if len(client.hash) != len(hash) {
		return fmt.Errorf("hash invalid")
	}

	for i := 0; i < len(client.hash); i++ {
		client.hash[i] = byte(hash[i])
	}

	return nil
}
