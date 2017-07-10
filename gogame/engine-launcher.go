package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/orolol/gogame/utils"
	"github.com/zeromq/goczmq"
)

//InitializePlayerDefaultValue init player
func InitializePlayerDefaultValue(acc utils.Account) utils.PlayerInGame {
	army := utils.PlayerArmy{
		NbSoldier:  100000,
		NbLigtTank: 100,
		NbHvyTank:  50,
		NbArt:      50,
		NbAirSup:   0,
		NbAirBomb:  0,
		Morale:     100,
		Quality:    100}

	economy := utils.PlayerEconomy{
		Money:   100000000,
		TaxRate: 5}

	civilian := utils.PlayerCivilian{
		NbTotalCivil:       60000000,
		NbManpower:         600000,
		NbHeavyTankFactory: 20,
		NbLightTankFactory: 20,
		NbCivilianFactory:  20}

	policy := utils.PlayerModifierPolicy{
		RecruitmentPolicy:  5,
		ManpowerSizePolicy: 1,
		ArtOnFactory:       false,
		BuildHvyTankFac:    true,
		BuildLgtTankFac:    true}

	var player = utils.PlayerInGame{
		PlayerID:       int(acc.ID),
		ModifierPolicy: policy,
		Army:           army,
		Nick:           acc.Name,
		Economy:        economy,
		Civilian:       civilian}

	return player
}

//PlayerAction player action
type PlayerAction func(player *utils.PlayerInGame, value float32)

//PASetRecruitementPolicy change recruitement policy to the value
func PASetRecruitementPolicy(player *utils.PlayerInGame, value float32) {
	qualityChange := player.Army.Quality - (100 - (1 / value))
	fmt.Println("QUALITY CHANGE ", qualityChange)
	player.Army.Quality -= value
	player.ModifierPolicy.RecruitmentPolicy = value
}
func setTaxRatePolicy(player *utils.PlayerInGame, value float32) {
	player.Economy.TaxRate = value
}
func setConscPolicy(player *utils.PlayerInGame, value float32) {
	player.Civilian.NbManpower -= player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.001
	player.Civilian.NbTotalCivil += player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.001
	player.ModifierPolicy.ManpowerSizePolicy = value
	player.Civilian.NbManpower += player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.001
	player.Civilian.NbTotalCivil -= player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.001
}
func setBuildLgtTank(player *utils.PlayerInGame, value float32) {
	if value == 1.0 {
		player.ModifierPolicy.BuildLgtTankFac = true
	} else {
		player.ModifierPolicy.BuildLgtTankFac = false
	}
}
func setBuildHvyTank(player *utils.PlayerInGame, value float32) {
	if value == 1.0 {
		player.ModifierPolicy.BuildHvyTankFac = true
	} else {
		player.ModifierPolicy.BuildHvyTankFac = false
	}
}

func actionCivConvertFactoryToLightTankFact(player *utils.PlayerInGame, value float32) {
	if player.Civilian.NbCivilianFactory > value {
		player.Civilian.NbCivilianFactory -= value
		player.Civilian.NbLightTankFactory += value
		player.Economy.Money -= 1000000 * value
	}
}
func actionCivConvertFactoryToHvyTankFact(player *utils.PlayerInGame, value float32) {
	if player.Civilian.NbCivilianFactory > value {
		player.Civilian.NbCivilianFactory -= value
		player.Civilian.NbHeavyTankFactory += value
		player.Economy.Money -= 1000000 * value
	}
}

func actionWarPropaganda(player *utils.PlayerInGame, value float32) {
	player.Economy.Money -= 10000000 * value
	player.Army.Morale += 15

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
		"setPopRecPolicy":                        PASetRecruitementPolicy,
		"setTaxRatePolicy":                       setTaxRatePolicy,
		"setBuildLgtTank":                        setBuildLgtTank,
		"setBuildHvyTank":                        setBuildHvyTank,
		"setConscPolicy":                         setConscPolicy,
		"actionWarPropaganda":                    actionWarPropaganda,
		"actionCivConvertFactoryToHvyTankFact":   actionCivConvertFactoryToHvyTankFact,
		"actionCivConvertFactoryToLightTankFact": actionCivConvertFactoryToLightTankFact,
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

		player2 = utils.AlgoDamageRepartition(player2, utils.AlgoDamageDealt(preFightP1))
		player1 = utils.AlgoDamageRepartition(player1, utils.AlgoDamageDealt(preFightP2))

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
