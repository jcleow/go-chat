package main

import (
	"fmt"
	"net/http"
	"github.com/jcleow/go-chat/pkg/websocket"
)

// Define websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request){
	fmt.Println("WebSocket endpoint hit")

	//upgrade connection to a websocket connection
	//https://stackoverflow.com/questions/50204967/what-is-websocket-upgrader-exactly
	conn,err := websocket.Upgrade(w, r)
	if err != nil{
		fmt.Fprintf(w,"%+V\n",err)
	}

	pool.Register <- client
	client.Read()
}


func setupRoutes(){	
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request){
		serveWs(pool, w, r)
	})
}

func main(){
	setupRoutes()
	fmt.Println("Chat App v0.01")
	http.ListenAndServe(":8080",nil)
}

