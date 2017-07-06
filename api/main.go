package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/orolol/utils"
)

var addr = flag.String("addr", ":5001", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 418)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func GameStateRouter(hub *Hub, queueGameState chan [][]byte) {
	for msg := range queueGameState {
		var gs utils.Game
		json.Unmarshal(msg[2], &gs)
		fmt.Println("GAME STATE RECIEVE: ", gs)

		for client := range hub.clients {
			if client.GameID == gs.GameID {
				fmt.Println("GAMEID FOUND ")
				w, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Println("ERROR ", err)
				} else {
					w.Write(msg[2])
				}
			} else if client.PlayerID == gs.ListPlayers[0].PlayerID || client.PlayerID == gs.ListPlayers[1].PlayerID {
				fmt.Println("PLAYERID FOUND ")
				client.GameID = gs.GameID
				w, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Println("ERROR ", err)
				} else {
					w.Write(msg[2])
				}
			}
		}
	}
}

func goSocket() {
	flag.Parse()
	hub := newHub()
	var queueGameState = make(chan [][]byte)
	go ZMQReader(queueGameState)

	go hub.run()
	go GameStateRouter(hub, queueGameState)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
		fmt.Println(hub)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&utils.Account{})
	db.AutoMigrate(&utils.Token{})

	go matchmaking()
	go goSocket()

	router := NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*", "localhost"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
