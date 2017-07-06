package main

import (
	"encoding/json"
	"fmt"

	"github.com/orolol/utils"
)

func matchmaking() {

	poolPendingPlayer := []utils.Account{}

	for account := range matchmakingQueue {
		fmt.Println("Recieving new match making msg ", account)
		poolPendingPlayer = append(poolPendingPlayer, account)
		if len(poolPendingPlayer) == 2 {
			CreateGame(poolPendingPlayer[0], poolPendingPlayer[1])
			poolPendingPlayer = []utils.Account{}
		}
	}
}

func CreateGame(p1, p2 utils.Account) {

	gc := utils.GameConf{GameType: "test", NbPlayers: 2, Players: []utils.Account{p1, p2}}

	jsonMsg, err := json.Marshal(gc)
	fmt.Println(string(jsonMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SEND CREATE GAME!")
	ZMQPusher.SendChan <- [][]byte{[]byte("CREATE"), []byte(jsonMsg)}
}

// function returns a channel
func createMatchMakingChan() chan utils.Account {
	b := make(chan utils.Account)
	return b
}

var matchmakingQueue = createMatchMakingChan()
