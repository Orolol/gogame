package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/zeromq/goczmq"
)

//InitializePlayerDefaultValue init player
func InitializePlayerDefaultValue(idPlayer int) PlayerInGame {
	army := PlayerArmy{
		NbSoldier:  1000,
		NbLigtTank: 100,
		NbHvyTank:  50,
		NbArt:      50,
		NbAirSup:   0,
		NbAirBomb:  0}

	policy := PlayerModifierPolicy{
		RecruitmentPolicy: 5}

	var player = PlayerInGame{
		PlayerID:       idPlayer,
		ModifierPolicy: policy,
		Army:           army,
		nick:           "Player " + strconv.Itoa(idPlayer),
		NbPop:          10000}

	return player
}

//PlayerAction player action
type PlayerAction func(player *PlayerInGame, value float32)

//PASetRecruitementPolicy change recruitement policy to the value
func PASetRecruitementPolicy(player *PlayerInGame, value float32) {
	fmt.Println("Change Rec policy to ", value)
	player.ModifierPolicy.RecruitmentPolicy = value
}

func createGame(conf GameConf, queue chan GameMsg) Game {
	var gameID = uuid.New()
	var mockP1 = InitializePlayerDefaultValue(conf.PlayerIDS[0])
	var mockP2 = InitializePlayerDefaultValue(conf.PlayerIDS[1])
	var listPlayer = []PlayerInGame{mockP1, mockP2}
	var game = Game{
		GameID:      gameID,
		CurrentTurn: 0,
		ListPlayers: listPlayer,
		Conf:        conf}

	return game
}

func GameEvent(queue chan GameMsg, game Game, player1, player2 *PlayerInGame) {
	ActionMapping := map[string]interface{}{
		"PASetRecruitementPolicy": PASetRecruitementPolicy,
	}
	for msg := range queue {
		if player1.PlayerID == msg.PlayerID {
			ActionMapping[msg.Action].(func(*PlayerInGame, float32))(player1, msg.Value)
		} else {
			ActionMapping[msg.Action].(func(*PlayerInGame, float32))(player2, msg.Value)
		}
		fmt.Println(game.ListPlayers)

	}
}

func runGame(game Game, queue chan GameMsg, queueGameOut chan Game) {
	var player1, player2 *PlayerInGame
	player1 = &game.ListPlayers[0]
	player2 = &game.ListPlayers[1]

	go GameEvent(queue, game, player1, player2)

	fmt.Println("Start game ", player1.nick, " vs ", player2.nick)

	for game.CurrentTurn < 999 {
		timer1 := time.NewTimer(time.Second / 4)
		//Resolve combat
		var preFightP1 = player1
		var preFightP2 = player2

		player2 = AlgoDamageRepartition(player2, AlgoDamageDealt(preFightP1))
		player1 = AlgoDamageRepartition(player1, AlgoDamageDealt(preFightP2))

		player2.Army.NbSoldier += AlgoReinforcement(player2)
		player1.Army.NbSoldier += AlgoReinforcement(player1)

		player2.NbPop -= AlgoReinforcement(player2)
		player1.NbPop -= AlgoReinforcement(player1)

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

		fmt.Println("Current State ", game)
	}
	fmt.Println("End game")
}

func GameManagerF(queueGameOut chan Game, queueCreation chan [][]byte) {

	var GameList = make(map[uuid.UUID]chan GameMsg)

	var gcM = GameConf{GameType: "FIRST MOCK", NbPlayers: 2, PlayerIDS: []int{999, 666}}
	queueGameInc := make(chan GameMsg, 100)
	game := createGame(gcM, queueGameInc)
	go runGame(game, queueGameInc, queueGameOut)
	uuid, err := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if err != nil {
		fmt.Println("err")
	}
	GameList[uuid] = queueGameInc

	for msg := range queueCreation {
		switch string(msg[1]) {
		case "CREATE":
			var gc GameConf
			json.Unmarshal(msg[2], &gc)
			fmt.Println("GAME CREATE : ", gc)
			fmt.Printf("Launch Game from this msg ", gc)
			queueGameInc := make(chan GameMsg, 100)
			game := createGame(gc, queueGameInc)
			go runGame(game, queueGameInc, queueGameOut)
			GameList[game.GameID] = queueGameInc
			fmt.Println("CURRENT LIST OF GAME ", GameList)
		case "MSG":
			var gs GameMsg
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
func ZMQPusherMockMSG() *goczmq.Channeler {
	fmt.Printf("Init Pusher")
	push := goczmq.NewDealerChanneler("tcp://127.0.0.1:31337")

	return push
}

func FromChanToZMQ(queue chan Game) {
	pushSock := ZMQPusher()
	for msg := range queue {
		fmt.Println("Recieving new game state")
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("fail :(")
			fmt.Println(err)
		}
		pushSock.SendChan <- [][]byte{[]byte(msg.GameID.String()), []byte(jsonMsg)}
	}
}

func main() {
	fmt.Printf("Enter Main")
	queueGameOut := make(chan Game, 100)
	queueGameIn := make(chan [][]byte, 100)

	go GameManagerF(queueGameOut, queueGameIn)
	go ZMQReader(queueGameIn)
	go FromChanToZMQ(queueGameOut)

	// pushSockMsg := ZMQPusherMockMSG()
	// gConf := GameConf{GameType: "test", NbPlayers: 2, PlayerIDS: []int{1, 2}}
	// jsonMsg, err := json.Marshal(gConf)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("SEND CREATE GAME!")
	// pushSockMsg.SendChan <- [][]byte{[]byte("CREATE"), []byte(jsonMsg)}
	//
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
