package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)
	// sockets
	server.OnConnect("connection", func(so socketio.Conn) error {
		log.Println("A new user connected")
		so.Join("chat_room")

		so.On("chat message", func(msg string) {
			log.Println("emit: ", so.Emit("chat message", msg))
			so.BroadcastTo("chat_room", "chat message", msg)
		})
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
