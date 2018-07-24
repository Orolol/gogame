package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/orolol/gogame/utils"
)

var poolPendingPlayerAI = []utils.Account{}
var poolPendingPlayer = []utils.Account{}

func matchmaking() {

	for account := range matchmakingQueue {
		fmt.Println("Recieving new match making msg ", account)
		fmt.Println("1 CURRENT MACTHMAKING QUEUE ", poolPendingPlayer)
		var isOk = true
		for _, p := range poolPendingPlayer {
			if p.ID == account.ID {
				fmt.Println("Already in queue")
				isOk = false
			}
		}
		if isOk {
			poolPendingPlayer = append(poolPendingPlayer, account)
			if len(poolPendingPlayer)%2 == 0 {
				time.Sleep(100 * time.Millisecond)
				CreateGame(poolPendingPlayer[0], poolPendingPlayer[1])
				poolPendingPlayer = []utils.Account{}
			}
			fmt.Println("AFTER CURRENT MACTHMAKING QUEUE ", poolPendingPlayer)
		}

	}

}

func leaveMatchmaking() {
	for account := range leaveMatchmakingQueue {
		fmt.Println("Recieving leaving ", account)
		fmt.Println("CURRENT MACTHMAKING QUEUE ", poolPendingPlayer)
		var toremove int = 0
		var isremove = false
		for i, p := range poolPendingPlayer {
			if p.ID == account.ID {
				toremove = i
				isremove = true
			}
		}
		if isremove {
			poolPendingPlayer = remove(poolPendingPlayer, toremove)
		}
		fmt.Println("CURRENT MACTHMAKING QUEUE ", poolPendingPlayer)

	}
}

func remove(s []utils.Account, i int) []utils.Account {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func matchmakingAi() {

	for account := range matchmakingAiQueue {
		fmt.Println("Recieving new AI MATCH ", account)
		fmt.Println("1 CURRENT MACTHMAKING QUEUE ", poolPendingPlayerAI)
		var isOk = true
		for _, p := range poolPendingPlayerAI {
			if p.ID == account.ID {
				fmt.Println("Already in queue")
				isOk = false
			}
		}
		if isOk {
			poolPendingPlayerAI = append(poolPendingPlayerAI, account)
			CreateAiGame(poolPendingPlayerAI[0])
			poolPendingPlayerAI = []utils.Account{}
			fmt.Println("AFTER CURRENT MACTHMAKING QUEUE ", poolPendingPlayerAI)
		}

	}
}

func CreateGame(p1, p2 utils.Account) {

	gc := utils.GameConf{GameType: "pvp", NbPlayers: 2, Players: []utils.Account{p1, p2}}

	jsonMsg, err := json.Marshal(gc)
	fmt.Println(string(jsonMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SEND CREATE GAME!")
	ZMQPusher.SendChan <- [][]byte{[]byte("CREATE"), []byte(jsonMsg)}
}

func CreateAiGame(p1 utils.Account) {

	var p2 = utils.Account{
		Login:           "AI",
		ELO:             1500,
		SelectedCountry: "Germany",
	}

	gc := utils.GameConf{GameType: "AI", NbPlayers: 2, Players: []utils.Account{p1, p2}}

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

// function returns a channel
func createMatchMakingChanAi() chan utils.Account {
	b := make(chan utils.Account)
	return b
}

var matchmakingQueue = createMatchMakingChan()
var leaveMatchmakingQueue = createMatchMakingChan()
var matchmakingAiQueue = createMatchMakingChanAi()
