package http

import (
	"net/http"
	"usercontrol/src/server/tcp"
)

func Run() {
	http.HandleFunc("/v1/setup/start", func(w http.ResponseWriter, r *http.Request) {
		tcp.SendCommand([]byte("start-setup"))

		w.WriteHeader(200)
	})

	http.HandleFunc("/v1/setup/end", func(w http.ResponseWriter, r *http.Request) {
		tcp.SendCommand([]byte("end-setup"))

		w.WriteHeader(200)
	})

	http.ListenAndServe("0.0.0.0:8889", nil)
}
