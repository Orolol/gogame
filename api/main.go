package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/orolol/gogame/utils"
)

var addr = flag.String("addr", ":5001", "http service address")

var ConnexionString = "root:@/gogame?charset=utf8&parseTime=True&loc=Local"

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
					w.Write(msg[2])
				}
			}
		}

		if gs.State == "End" {
			fmt.Println("END GAME")
			if gs.Conf.GameType != "AI" {
				db, _ := gorm.Open("mysql", ConnexionString)
				delete(onGoingGames, gs.GameID)
				var winner utils.Account
				var loser utils.Account
				var gh utils.GameHistory
				db.Where("ID = ? ", gs.Winner.PlayerID).First(&winner)
				db.Where("ID = ? ", gs.Loser.PlayerID).First(&loser)

				gh.WinnerID = winner.ID
				// gh.WinnerNick = winner.Name
				gh.LoserID = loser.ID
				// gh.LoserNick = loser.Name
				gh.GameID = gs.GameID
				gh.ELODiff = 15

				db.Create(&gh)

				winner.ELO += 15
				loser.ELO -= 15
				db.Save(winner)
				db.Save(loser)
			} else {
				fmt.Println(1)
				db, _ := gorm.Open("mysql", ConnexionString)
				delete(onGoingGames, gs.GameID)
				var winner utils.Account
				var loser utils.Account
				var gh utils.GameHistory
				fmt.Println(2)
				if gs.Winner.PlayerID != 0 {
					db.Where("ID = ? ", gs.Winner.PlayerID).First(&winner)

				} else {
					db.Where("ID = ? ", gs.Loser.PlayerID).First(&loser)
				}
				fmt.Println(3)
				gh.WinnerID = winner.ID
				// gh.WinnerNick = winner.Name
				gh.LoserID = loser.ID
				// gh.LoserNick = loser.Name
				gh.GameID = gs.GameID
				gh.ELODiff = 0
				fmt.Println(4, gh)
				db.Create(&gh)
				fmt.Println(5)
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
						w.Write(msg[2])
					}
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
	// http.HandleFunc("/", serveHome)
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

	db, err := gorm.Open("mysql", ConnexionString)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&utils.Account{})
	db.AutoMigrate(&utils.GameHistory{})
	db.AutoMigrate(&utils.Token{})

	utils.SetBaseValueDB()

	go matchmaking()
	go matchmakingAi()
	go goSocket()
	initRoutes()
	// router := NewRouter()
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// originsOk := handlers.AllowedOrigins([]string{"*", "localhost"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
