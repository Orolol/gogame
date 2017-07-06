package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/orolol/utils"
	"github.com/zeromq/goczmq"
)

//InitializePlayerDefaultValue init player
func InitializePlayerDefaultValue(acc utils.Account) utils.PlayerInGame {
	army := utils.PlayerArmy{
		NbSoldier:  1000,
		NbLigtTank: 100,
		NbHvyTank:  50,
		NbArt:      50,
		NbAirSup:   0,
		NbAirBomb:  0}

	policy := utils.PlayerModifierPolicy{
		RecruitmentPolicy: 5}

	var player = utils.PlayerInGame{
		PlayerID:       int(acc.ID),
		ModifierPolicy: policy,
		Army:           army,
		Nick:           acc.Name,
		NbPop:          10000}

	return player
}

//PlayerAction player action
type PlayerAction func(player *utils.PlayerInGame, value float32)

//PASetRecruitementPolicy change recruitement policy to the value
func PASetRecruitementPolicy(player *utils.PlayerInGame, value float32) {
	fmt.Println("Change Rec policy to ", value)
	player.ModifierPolicy.RecruitmentPolicy = value
}

func createGame(conf utils.GameConf, queue chan utils.GameMsg) utils.Game {
	var gameID = uuid.New()
	var mockP1 = InitializePlayerDefaultValue(conf.Players[0])
	var mockP2 = InitializePlayerDefaultValue(conf.Players[1])
	var listPlayer = []utils.PlayerInGame{mockP1, mockP2}
	var game = utils.Game{
		GameID:      gameID,
		CurrentTurn: 0,
		ListPlayers: listPlayer,
		Conf:        conf}

	return game
}

func GameEvent(queue chan utils.GameMsg, game utils.Game, player1, player2 *utils.PlayerInGame) {
	ActionMapping := map[string]interface{}{
		"PASetRecruitementPolicy": PASetRecruitementPolicy,
	}
	for msg := range queue {
		if player1.PlayerID == msg.PlayerID {
			ActionMapping[msg.Action].(func(*utils.PlayerInGame, float32))(player1, msg.Value)
		} else {
			ActionMapping[msg.Action].(func(*utils.PlayerInGame, float32))(player2, msg.Value)
		}
		fmt.Println(game.ListPlayers)

	}
}

func runGame(game utils.Game, queue chan utils.GameMsg, queueGameOut chan utils.Game) {
	var player1, player2 *utils.PlayerInGame
	player1 = &game.ListPlayers[0]
	player2 = &game.ListPlayers[1]

	go GameEvent(queue, game, player1, player2)

	fmt.Println("Start game ", player1.Nick, " vs ", player2.Nick)

	for game.CurrentTurn < 999 {
		timer1 := time.NewTimer(time.Second / 4)
		//Resolve combat
		var preFightP1 = player1
		var preFightP2 = player2

		player2 = utils.AlgoDamageRepartition(player2, utils.AlgoDamageDealt(preFightP1))
		player1 = utils.AlgoDamageRepartition(player1, utils.AlgoDamageDealt(preFightP2))

		player2.Army.NbSoldier += utils.AlgoReinforcement(player2)
		player1.Army.NbSoldier += utils.AlgoReinforcement(player1)

		player2.NbPop -= utils.AlgoReinforcement(player2)
		player1.NbPop -= utils.AlgoReinforcement(player1)

		if player1.NbPop <= 0 || player1.Army.NbSoldier <= 0 {
			fmt.Println("PLAYER 2 WIN ! ", game)
			queueGameOut <- game
			break
		}
		if player2.NbPop <= 0 || player2.Army.NbSoldier <= 0 {
			fmt.Println("PLAYER 1 WIN ! ", game)
			queueGameOut <- game
			break
		}
		<-timer1.C
		game.CurrentTurn++
		queueGameOut <- game

	}
	fmt.Println("End game")
}

func GameManagerF(queueGameOut chan utils.Game, queueCreation chan [][]byte) {

	var GameList = make(map[uuid.UUID]chan utils.GameMsg)

	for msg := range queueCreation {
		switch string(msg[1]) {
		case "CREATE":
			var gc utils.GameConf
			json.Unmarshal(msg[2], &gc)
			fmt.Println("GAME CREATE : ", gc)
			fmt.Println("Launch Game from this msg ", gc)
			queueGameInc := make(chan utils.GameMsg, 100)
			game := createGame(gc, queueGameInc)
			go runGame(game, queueGameInc, queueGameOut)
			GameList[game.GameID] = queueGameInc
			fmt.Println("CURRENT LIST OF GAME ", GameList)
		case "MSG":
			var gs utils.GameMsg
			json.Unmarshal(msg[2], &gs)
			fmt.Println("GameMSG : ", gs)
			if val, ok := GameList[gs.GameID]; ok {
				val <- gs
			} else {
				fmt.Println("Game not found ", gs.GameID)
			}

		}
	}

}

func ZMQReader(queueCreation chan [][]byte) {
	fmt.Printf("Init Reader")
	pull := goczmq.NewRouterChanneler("tcp://127.0.0.1:31337")
	for msg := range pull.RecvChan {
		fmt.Println("Recieving new game msg in ZMQ !! TYPE : ", string(msg[1]))
		queueCreation <- msg
	}
}
func ZMQPusher() *goczmq.Channeler {
	fmt.Printf("Init Pusher")
	push := goczmq.NewDealerChanneler("tcp://127.0.0.1:31338")

	return push
}

func FromChanToZMQ(queue chan utils.Game) {
	pushSock := ZMQPusher()
	for msg := range queue {
		fmt.Println("Recieving new game state")
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("fail :(")
			fmt.Println(err)
		}

		pushSock.SendChan <- [][]byte{[]byte(msg.GameID.String()), []byte(jsonMsg)}
		fmt.Println("SENT : ", msg)
	}
}

func main() {
	fmt.Printf("Enter Main")
	queueGameOut := make(chan utils.Game)
	queueGameIn := make(chan [][]byte)

	go GameManagerF(queueGameOut, queueGameIn)
	go ZMQReader(queueGameIn)
	go FromChanToZMQ(queueGameOut)

	// pushSockMsg := ZMQPusherMockMSG()
	// gConf := utils.GameConf{GameType: "test", NbPlayers: 2, PlayerIDS: []int{1, 2}}
	// jsonMsg, err := json.Marshal(gConf)
	// fmt.Println(string(jsonMsg))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("SEND CREATE GAME!")
	// pushSockMsg.SendChan <- [][]byte{[]byte("CREATE"), []byte(jsonMsg)}

	// gConf = GameConf{GameType: "test", NbPlayers: 2, PlayerIDS: []int{8, 9}}
	// jsonMsg, err = json.Marshal(gConf)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("SEND CREATE GAME!")
	// pushSockMsg.SendChan <- [][]byte{[]byte("CREATE"), []byte(jsonMsg)}

	//MOCK TO SEND EVENT
	// time.Sleep(2 * time.Second)
	// lGmsg := GameMsg{Action: "PASetRecruitementPolicy", PlayerID: 999, Text: "Change rec value to 5", Value: 15.0}
	// jsonMsg, err = json.Marshal(lGmsg)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// var gsa GameMsg
	// json.Unmarshal(jsonMsg, &gsa)
	// pushSockMsg.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}
	// time.Sleep(1 * time.Second)
	// lGmsg = GameMsg{Action: "PASetRecruitementPolicy", PlayerID: 1, Text: "Change rec value to 15", Value: 15}
	// jsonMsg, err = json.Marshal(lGmsg)
	// if err != nil {
	// 	fmt.Println("fail :(")
	// 	fmt.Println(err)
	// }
	// fmt.Println("SEND EVENT MOCK !")
	// fmt.Println(jsonMsg)
	// pushSockMsg.SendChan <- [][]byte{[]byte(GID.String()), []byte(jsonMsg)}
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
