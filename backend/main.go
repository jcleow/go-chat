package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// Defining an upgrader
// this will require a read and write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
  	// this will allow us to make requests from our React
  	// development server to here.
  	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool{ return true},
}

// Defining a reader which will listen for new messages
// being sent to our WebSocket endpoint
func reader(conn *websocket.Conn){
	for{
	// read in a message
	// p is []byte and messageType is an int with 
	// value websocket.BinaryMessage or websocket.TextMessage
	// https://pkg.go.dev/github.com/gorilla/websocket
		messageType, p, err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
	// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil{
			log.Println(err)
			return
		}
	}
}

// Define websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)

	//upgrade connection to a websocket connection
	//https://stackoverflow.com/questions/50204967/what-is-websocket-upgrader-exactly
	ws,err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
		reader(ws)
}


func setupRoutes(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Simple Server")
	})
	// map our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws",serveWs)
}

func main(){
	setupRoutes()
	fmt.Println("Chat App v0.01")
	http.ListenAndServe(":8080",nil)
}

