package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		/* upgrade HTTP connection to a WebSocket connection */
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				/* read message from connection */
				mType, msg, _ := conn.ReadMessage()
				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()
				fmt.Println(string(msg))
			}
		}(conn)
	})

	http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			ch := time.Tick(5 * time.Second)
			/* every 5 seconds send some data */
			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "bin3xis",
					FirstName: "Alexis",
					LastName:  "Rodriguez",
				})
			}
		}(conn)
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		/* goroutine checking for a close connection */
		go func(conn *websocket.Conn) {
			for {
				_, _, err := conn.ReadMessage()
				/* if error is not nil, close WebSocket*/
				if err != nil {
					conn.Close()
				}
			}
		}(conn)

		/* goroutine sending data */
		go func(conn *websocket.Conn) {
			ch := time.Tick(5 * time.Second)
			/* every 5 seconds send some data */
			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "bin3xis",
					FirstName: "Alexis",
					LastName:  "Rodriguez",
				})
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}
