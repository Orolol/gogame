package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/orolol/gogame/utils"
	"github.com/zeromq/goczmq"
)

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
		"setPopRecPolicy":                        PASetRecruitementPolicy,
		"setTaxRatePolicy":                       setTaxRatePolicy,
		"setBuildLgtTank":                        setBuildLgtTank,
		"setBuildHvyTank":                        setBuildHvyTank,
		"setConscPolicy":                         setConscPolicy,
		"actionWarPropaganda":                    actionWarPropaganda,
		"buyForeignTanks":                        buyForeignTanks,
		"actionCivConvertFactoryToHvyTankFact":   actionCivConvertFactoryToHvyTankFact,
		"actionCivConvertFactoryToLightTankFact": actionCivConvertFactoryToLightTankFact,
		"technoIndusT1N3":                        technoIndusT1N3,
		"technoIndusT1N2":                        technoIndusT1N2,
		"technoIndusT1N1":                        technoIndusT1N1,
		"technoIndusT2N3":                        technoIndusT2N3,
		"technoIndusT2N2":                        technoIndusT2N2,
		"technoIndusT2N1":                        technoIndusT2N1,
		"technoMilT1N3":                          technoMilT1N3,
		"technoMilT1N2":                          technoMilT1N2,
		"technoMilT1N1":                          technoMilT1N1,
		"technoMilT2N3":                          technoMilT2N3,
		"technoMilT2N2":                          technoMilT2N2,
		"technoMilT2N1":                          technoMilT2N1,
	}
	for msg := range queue {
		if player1.PlayerID == msg.PlayerID {
			if len(msg.Effects) > 0 {
				genericApplyEffect(player1, msg.Effects)
			}
			ActionMapping[msg.Action].(func(*utils.PlayerInGame, map[string]float32, []utils.Effect))(player1, msg.Value, msg.Effects)
		} else {
			if len(msg.Effects) > 0 {
				genericApplyEffect(player2, msg.Effects)
			}
			ActionMapping[msg.Action].(func(*utils.PlayerInGame, map[string]float32, []utils.Effect))(player2, msg.Value, msg.Effects)
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
	game.State = "Running"
	queueGameOut <- game
	queueGameOut <- game
	time.Sleep(5 * time.Second)
	for game.CurrentTurn < 9999 {
		timer1 := time.NewTimer(time.Second)
		game.CurrentTurn++
		//Resolve combat
		var preFightP1 = player1
		var preFightP2 = player2
		if game.CurrentTurn > 30 {
			player2 = utils.AlgoDamageRepartition(player2, utils.AlgoDamageDealt(preFightP1))
			player1 = utils.AlgoDamageRepartition(player1, utils.AlgoDamageDealt(preFightP2))
		}

		if player1.Army.NbSoldier <= 0 {
			fmt.Println("P2 WIN")
			game.State = "End"
			game.Winner = game.ListPlayers[1]
			game.Loser = game.ListPlayers[0]
			fmt.Println("SEND ", game)
			queueGameOut <- game
			queueGameOut <- game
			break
		} else if player2.Army.NbSoldier <= 0 {
			fmt.Println("P1 WIN")
			game.State = "End"
			game.Winner = game.ListPlayers[0]
			game.Loser = game.ListPlayers[1]
			fmt.Println("SEND ", game)
			queueGameOut <- game
			queueGameOut <- game
			break
		} else {
			player2 = utils.AlgoReinforcement(player2)
			player1 = utils.AlgoReinforcement(player1)

			if player2.Civilian.NbManpower < 0 {
				player2.Civilian.NbManpower = 0.0
			}
			if player1.Civilian.NbManpower < 0 {
				player1.Civilian.NbManpower = 0.0
			}

			player1 = utils.AlgoEconomicEndTurn(player1)
			player2 = utils.AlgoEconomicEndTurn(player2)

			fmt.Println("NO WIN")
			<-timer1.C
			fmt.Println("SEND ", game)
			queueGameOut <- game
			queueGameOut <- game
		}

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
		fmt.Println("Read Game from Queue and send to ZMQ")
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

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
