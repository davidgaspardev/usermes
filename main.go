package main

import (
	"usercontrol/src/server/http"
	"usercontrol/src/server/tcp"
)

func main() {
	go tcp.Run()
	http.Run()
}
