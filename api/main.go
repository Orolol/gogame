package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/orolol/gogame/utils"
)

var addr = flag.String("addr", ":5001", "http service address")

var onGoingGames = make(map[uuid.UUID]*utils.Game)

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

		if gs.State == "End" {
			db, _ := gorm.Open("sqlite3", "test.db")
			delete(onGoingGames, gs.GameID)
			var winner utils.Account
			var loser utils.Account
			db.Where("ID = ? ", gs.Winner.PlayerID).First(&winner)
			db.Where("ID = ? ", gs.Loser.PlayerID).First(&loser)

			winner.ELO += 15
			loser.ELO -= 15
			db.Save(winner)
			db.Save(loser)
		}

		for client := range hub.clients {
			if client.GameID == gs.GameID {
				onGoingGames[gs.GameID] = &gs
				w, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Println("ERROR ", err)
				} else {
					w.Write(msg[2])
				}
			} else if client.PlayerID == gs.ListPlayers[0].PlayerID || client.PlayerID == gs.ListPlayers[1].PlayerID {
				client.GameID = gs.GameID
				onGoingGames[gs.GameID] = &gs
				w, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Println("ERROR ", err)
				} else {
					fmt.Println("WRITE !")
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
	db.AutoMigrate(&utils.Policy{})
	db.AutoMigrate(&utils.PlayerActionOrder{})
	db.AutoMigrate(&utils.Technology{})

	utils.SetBaseValueDB()

	go matchmaking()
	go goSocket()

	router := NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*", "localhost"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
