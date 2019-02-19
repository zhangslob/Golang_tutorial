package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ws)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			return
		}
		log.Printf("msg: %s", string(msg))
	}
}
