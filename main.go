package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/wss", func(rw http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		c, err := upgrader.Upgrade(rw, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			returnMessage := []byte(fmt.Sprintf("received message `%s`", message))
			err = c.WriteMessage(mt, returnMessage)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})

	addr := flag.String("addr", "localhost:8181", "http service address")
	http.ListenAndServe(*addr, nil)
}
