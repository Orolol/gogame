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
	utils.SetBaseValueDB()
	var gameID = uuid.New()
	var mockP1 = InitializePlayerDefaultValue(conf.Players[0])
	var mockP2 = InitializePlayerDefaultValue(conf.Players[1])
	var listPlayer = []utils.PlayerInGame{mockP1, mockP2}
	var game = utils.Game{
		GameID:      gameID,
		CurrentTurn: 0,
		ListPlayers: listPlayer,
		IsWar:       false,
		Conf:        conf}

	return game
}

func GameEvent(queue chan utils.GameMsg, game *utils.Game, player1, player2 *utils.PlayerInGame) {
	// ActionMapping := map[string]interface{}{
	// 	// "actionWarPropaganda":                    actionWarPropaganda,
	// 	"buyForeignTanks": buyForeignTanks,
	// 	// "actionCivConvertFactoryToHvyTankFact":   actionCivConvertFactoryToHvyTankFact,
	// 	// "actionCivConvertFactoryToLightTankFact": actionCivConvertFactoryToLightTankFact,
	// 	"emergencyRecruitment": emergencyRecruitment,
	// 	"purgeSoldier":         purgeSoldier,
	// }
	// //fmt.Println("GAME EVENT ", game)
	// keys := make([]string, 0, len(ActionMapping))
	// for k := range ActionMapping {
	// 	keys = append(keys, k)
	// }

	for msg := range queue {
		var p *utils.PlayerInGame
		var opponent *utils.PlayerInGame

		//Determine the player
		if player1.PlayerID == msg.PlayerID {
			p = player1
			opponent = player2
		} else {
			p = player2
			opponent = player1
		}

		//Apply Effects, cost and special action
		if len(msg.Effects) > 0 {
			genericApplyEffect(p, opponent, msg.Effects, game)
		}
		if len(msg.Costs) > 0 {
			genericApplyCosts(p, msg.Costs)
		}
		if msg.Cooldown != 0 {
			var order = utils.PlayerLastOrders{
				OrderID:  msg.Action,
				Cooldown: msg.Cooldown + game.CurrentTurn,
			}
			p.LastOrders = append(p.LastOrders, order)
			//fmt.Println("MISE SOUS CD", msg.Cooldown, game.CurrentTurn)
		}
		// if utils.StringInSlice(msg.Action, keys) {
		// 	ActionMapping[msg.Action].(func(*utils.PlayerInGame, float32))(p, msg.Value)
		// }

		if msg.Type == "TECH" {
			p.Technologies = append(p.Technologies, msg.Action)
		} else if msg.Type == "POLICY" {
			var pol = utils.GetPolicy(msg.Action)
			var choosePol utils.PolicyValue
			for _, x := range pol.PossibleValue2 {
				//fmt.Println("FOUND POLICY ", x)
				if x.Value == msg.Value {
					choosePol = x
				}
			}
			for k, pl := range p.Policies {
				if pl.ActionName == msg.Action {
					p.Policies[k] = choosePol
				}

			}
		}
	}

}

func runGame(game utils.Game, queue chan utils.GameMsg, queueGameOut chan utils.Game) {
	var player1, player2 *utils.PlayerInGame
	player1 = &game.ListPlayers[0]
	player2 = &game.ListPlayers[1]

	go GameEvent(queue, &game, player1, player2)

	//fmt.Println("Start game ", player1.Nick, " vs ", player2.Nick)
	game.State = "Running"
	queueGameOut <- game
	queueGameOut <- game
	time.Sleep(time.Second)

	for _, e := range player1.Country.Effects {
		utils.ApplyEffect(player1, e, &game)
	}
	for _, e := range player2.Country.Effects {
		utils.ApplyEffect(player2, e, &game)
	}

	for game.CurrentTurn < 9999 {

		timer1 := time.NewTimer(time.Second * 1)
		game.CurrentTurn++
		fmt.Println("Turn : ", game.CurrentTurn)
		for i, rlen := 0, len(player1.CallbackEffects); i < rlen; i++ {
			j := i - (rlen - len(player1.CallbackEffects))
			var cb = player1.CallbackEffects[j]
			if utils.CheckConstraint(player1, cb.Constraints, nil, &game, 0) {
				for _, e := range cb.Effects {
					utils.ApplyEffect(player1, e, &game)
					if len(player1.CallbackEffects) > 1 {
						player1.CallbackEffects = append(player1.CallbackEffects[:j], player1.CallbackEffects[j+1:]...)
					} else {
						player1.CallbackEffects = player1.CallbackEffects[:0]
					}
				}
			}
		}
		for i, rlen := 0, len(player2.CallbackEffects); i < rlen; i++ {
			j := i - (rlen - len(player2.CallbackEffects))
			var cb = player2.CallbackEffects[j]
			if utils.CheckConstraint(player2, cb.Constraints, nil, &game, 0) {
				for _, e := range cb.Effects {
					utils.ApplyEffect(player2, e, &game)
					fmt.Println(1, j, player2.CallbackEffects)
				}
				if len(player1.CallbackEffects) > 1 {
					player2.CallbackEffects = append(player2.CallbackEffects[:j], player2.CallbackEffects[j+1:]...)
				} else {
					player2.CallbackEffects = player2.CallbackEffects[:0]
				}
			}
		}

		if game.Conf.GameType == "AI" {
			utils.RollAiAction(player2, &game)
		}
		//Event start turn
		player1 = utils.AlgoRollTurnEvent(player1, &game)
		player2 = utils.AlgoRollTurnEvent(player2, &game)

		//Resolve combat
		var preFightP1 = player1
		var preFightP2 = player2

		if game.CurrentTurn > 20 {
			game.IsWar = true
		}
		if game.IsWar {
			p1dmg := utils.AlgoDamageDealt(preFightP1)
			p2dmg := utils.AlgoDamageDealt(preFightP2)
			player2 = utils.AlgoDamageRepartition(player2, p1dmg)
			player1 = utils.AlgoDamageRepartition(player1, p2dmg)
			player1, player2 = utils.AlgoFullAerialPhase(player1, player2)
			player1, player2 = utils.AlgoTerritorryChange(player1, player2, p1dmg, p2dmg)

			p1FactDmg := utils.AlgoDamageDealtOnFactories(preFightP1)
			p2FactDmg := utils.AlgoDamageDealtOnFactories(preFightP2)

			player2 = utils.AlgoDamageRepartitionOnFactories(player2, p1FactDmg)
			player1 = utils.AlgoDamageRepartitionOnFactories(player1, p2FactDmg)

			// p1PopDmg := utils.AlgoDamageDealtOnPopulation(preFightP1)
			// p2PopDmg := utils.AlgoDamageDealtOnPopulation(preFightP2)
		}

		if player1.Territory.Surface <= 0 {
			game.State = "End"
			game.Winner = game.ListPlayers[1]
			game.Loser = game.ListPlayers[0]
			queueGameOut <- game
			break
		} else if player2.Territory.Surface <= 0 {
			game.State = "End"
			game.Winner = game.ListPlayers[0]
			game.Loser = game.ListPlayers[1]
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

			<-timer1.C
			queueGameOut <- game
			// queueGameOut <- game
		}

	}
}

func GameManagerF(queueGameOut chan utils.Game, queueCreation chan [][]byte) {

	var GameList = make(map[uuid.UUID]chan utils.GameMsg)

	for msg := range queueCreation {
		switch string(msg[1]) {
		case "CREATE":
			var gc utils.GameConf
			json.Unmarshal(msg[2], &gc)
			queueGameInc := make(chan utils.GameMsg, 100)
			game := createGame(gc, queueGameInc)
			go runGame(game, queueGameInc, queueGameOut)
			GameList[game.GameID] = queueGameInc
		case "MSG":
			var gs utils.GameMsg
			json.Unmarshal(msg[2], &gs)
			if val, ok := GameList[gs.GameID]; ok {
				val <- gs
			}

		}
	}

}

func ZMQReader(queueCreation chan [][]byte) {
	//fmt.Printf("Init Reader")
	pull := goczmq.NewRouterChanneler("tcp://127.0.0.1:31337")
	for msg := range pull.RecvChan {
		//fmt.Println("Recieving new game msg in ZMQ !! TYPE : ", string(msg[1]))
		queueCreation <- msg
	}
}
func ZMQPusher() *goczmq.Channeler {
	//fmt.Printf("Init Pusher")
	push := goczmq.NewDealerChanneler("tcp://127.0.0.1:31338")

	return push
}

func FromChanToZMQ(queue chan utils.Game) {
	pushSock := ZMQPusher()
	for msg := range queue {
		//fmt.Println("Read Game from Queue and send to ZMQ")
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			//
			//fmt.Println(err)
		}

		pushSock.SendChan <- [][]byte{[]byte(msg.GameID.String()), []byte(jsonMsg)}
		//fmt.Println("SENT : ", msg)
	}
}

func main() {
	//fmt.Printf("Enter Main")
	queueGameOut := make(chan utils.Game)
	queueGameIn := make(chan [][]byte)

	go GameManagerF(queueGameOut, queueGameIn)
	go ZMQReader(queueGameIn)
	go FromChanToZMQ(queueGameOut)

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
