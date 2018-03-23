package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/orolol/gogame/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")

}

func ChangePolicy(w http.ResponseWriter, r *http.Request) {
	var polChange utils.PolicyChange
	var pol utils.Policy
	var gMsg utils.GameMsg
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &polChange); err != nil {
		panic(err)
	}
	fmt.Println("CHANGE POLICY, ", polChange)

	pol = utils.GetPolicy(polChange.ID)
	gMsg.Action = pol.ActionName
	gMsg.GameID = polChange.GameID
	gMsg.PlayerID = polChange.PlayerID
	gMsg.Text = "Change pol"
	gMsg.Value = make(map[string]float32)
	gMsg.Value["value"] = polChange.Value

	jsonMsg, err := json.Marshal(gMsg)
	fmt.Println(string(jsonMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SEND Game MSG!")
	ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}

}

func GetTechnology(w http.ResponseWriter, r *http.Request) {
	var actionApi utils.PolicyChange
	var techno utils.Technology
	var gMsg utils.GameMsg

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &actionApi); err != nil {
		panic(err)
	}
	fmt.Println(onGoingGames)
	var isOkAction bool = true
	var game *utils.Game
	var ok bool
	var players int
	var tech int
	if game, ok = onGoingGames[actionApi.GameID]; ok {
		techno = utils.GetTechnolgy(actionApi.ID)
		for players = range game.ListPlayers {
			if game.ListPlayers[players].PlayerID == actionApi.PlayerID {
				fmt.Println("TECH CHECK", techno)
				for tech = range game.ListPlayers[players].Technologies {
					if game.ListPlayers[players].Technologies[tech] == techno.ActionName {
						fmt.Println("ALREADY GOT THE TECH")
						isOkAction = false
					}
				}
				if !utils.CheckConstraint(&game.ListPlayers[players], techno.Constraints, techno.Costs, game) {
					fmt.Println("CONSTRAINT FAIL")
					isOkAction = false
				} else {
					fmt.Println("CONSTRAINT OK")
				}
			}
		}

	}
	if isOkAction {
		game.ListPlayers[players].Technologies = append(game.ListPlayers[players].Technologies, techno.ActionName)
		gMsg.Action = techno.ActionName
		gMsg.GameID = actionApi.GameID
		gMsg.PlayerID = actionApi.PlayerID
		gMsg.Text = "Order"
		gMsg.Costs = techno.Costs
		gMsg.Effects = techno.Effects
		gMsg.Type = "TECH"
		fmt.Println("TECH UP !!!", gMsg, techno)
		jsonMsg, err := json.Marshal(gMsg)
		fmt.Println(string(jsonMsg))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(onGoingGames)
		ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}
	} else {
		fmt.Println("CANT TECH UP")
	}

}

func Actions(w http.ResponseWriter, r *http.Request) {
	var actionApi utils.PolicyChange
	var action utils.PlayerActionOrder
	var gMsg utils.GameMsg
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &actionApi); err != nil {
		panic(err)
	}
	fmt.Println(onGoingGames)
	var isOkAction bool = true
	if game, ok := onGoingGames[actionApi.GameID]; ok {
		fmt.Println("GOT THE GAME")
		action = utils.GetAction(actionApi.ID)
		for players := range game.ListPlayers {
			if game.ListPlayers[players].PlayerID == actionApi.PlayerID {
				fmt.Println("GOT THE PLayer")
				if len(game.ListPlayers[players].LastOrders) > 0 {
					for actions := range game.ListPlayers[players].LastOrders {
						if game.ListPlayers[players].LastOrders[actions].OrderID == action.ActionName {
							if game.ListPlayers[players].LastOrders[actions].Cooldown > game.CurrentTurn {
								fmt.Println("CD END ", game.ListPlayers[players].LastOrders[actions].Cooldown)
								isOkAction = false
							}
						}
					}
				}
				if !utils.CheckConstraint(&game.ListPlayers[players], action.Constraints, action.Costs, game) {
					fmt.Println("CONSTRAINT FAIL")
					isOkAction = false
				} else {
					fmt.Println("CONSTRAINT OK")
				}
			}
		}

	}
	if isOkAction {
		fmt.Println("OK ACTION")
		gMsg.Action = action.ActionName
		gMsg.GameID = actionApi.GameID
		gMsg.PlayerID = actionApi.PlayerID
		gMsg.Text = "Order"
		gMsg.Costs = action.Costs
		gMsg.Effects = action.Effects
		gMsg.Cooldown = action.Cooldown
		jsonMsg, err := json.Marshal(gMsg)
		fmt.Println(string(jsonMsg))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(onGoingGames)
		ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}
	} else {
		fmt.Println("ACTION ON CD")
	}

}

func JoinGame(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	fmt.Println("Seems like someone want to join a game ! ", r.Body)
	var acc utils.Account

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &acc); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.Where(&acc).First(&acc)

	var m = make(map[string]interface{})
	m["policies"] = getDefaultPolicies()
	m["actions"] = getDefaultActions()
	m["technology"] = getDefaultTech()
	m["events"] = getDefaultEvents()

	jsonMsg, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonMsg)

	matchmakingQueue <- acc
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var gc utils.GameMsg
	fmt.Println("Handle send msg")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &gc); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	jsonMsg, err := json.Marshal(gc)
	fmt.Println(string(jsonMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SEND GAME MSG!")
	ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
}
