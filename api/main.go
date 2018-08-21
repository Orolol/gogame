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
	"github.com/kortemy/elo-go"
	"github.com/orolol/gogame/utils"
)

var addr = flag.String("addr", ":5001", "http service address")

var ConnexionString = "gogame:See9ubeo#@/gogame?charset=utf8&parseTime=True&loc=Local"

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

			if gs.Conf.GameType != "AI" {
				fmt.Println("END GAME", gs.GameID)
				fmt.Println("onGoingGames", onGoingGames)
				db, _ := gorm.Open("mysql", ConnexionString)
				delete(onGoingGames, gs.GameID)
				var winner utils.Account
				var loser utils.Account
				var gh utils.GameHistory
				db.Where("ID = ? ", gs.Winner.PlayerID).First(&winner)
				db.Where("ID = ? ", gs.Loser.PlayerID).First(&loser)

				rankA := winner.ELO
				rankB := loser.ELO
				elo := elogo.NewElo()

				// Expected chance that A defeats B
				// use ExpectedScoreWithFactors(rankA, rankB, deviation) to use custom factor for this method
				elo.ExpectedScore(rankA, rankB) // 0.3599350001971149

				// Results for A in the outcome of A defeats B
				score := float64(1)                  // Use 1 in case A wins, 0 in case B wins, 0.5 in case of a draw)
				elo.RatingDelta(rankA, rankB, score) // 20
				elo.Rating(rankA, rankB, score)      // 1520
				outcomeA, outcomeB := elo.Outcome(rankA, rankB, score)

				gh.WinnerID = winner.ID
				// gh.WinnerNick = winner.Name
				gh.LoserID = loser.ID
				// gh.LoserNick = loser.Name
				gh.GameID = gs.GameID
				gh.ELODiff = outcomeA.Delta

				db.Create(&gh)

				winner.ELO = outcomeA.Rating
				loser.ELO = outcomeB.Rating
				db.Save(winner)
				db.Save(loser)
				fmt.Println("onGoingGames AFTER", onGoingGames)
			} else {
				db, _ := gorm.Open("mysql", ConnexionString)
				delete(onGoingGames, gs.GameID)
				var winner utils.Account
				var loser utils.Account
				var gh utils.GameHistory
				if gs.Winner.PlayerID != 0 {
					db.Where("ID = ? ", gs.Winner.PlayerID).First(&winner)

				} else {
					db.Where("ID = ? ", gs.Loser.PlayerID).First(&loser)
				}
				gh.WinnerID = winner.ID
				// gh.WinnerNick = winner.Name
				gh.LoserID = loser.ID
				// gh.LoserNick = loser.Name
				gh.GameID = gs.GameID
				gh.ELODiff = 0
				db.Create(&gh)
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
	go leaveMatchmaking()
	go goSocket()
	initRoutes()
	// router := NewRouter()
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// originsOk := handlers.AllowedOrigins([]string{"*", "localhost"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// log.Fatal(http.ListenAndServe(":8081", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
