package tcp

import (
	"net"

	"github.com/davidgaspardev/golog"
)

func Run() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		golog.Error("TCP", err)
	}

	defer listener.Close()
	golog.System("TCP", "started server in port: 8888")

	go getServer().run()

	for {
		conn, err := listener.Accept()
		if err != nil {
			golog.Error("TCP", err)
			continue
		}

		go server.connectClient(conn)
	}
}
