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
	var actionApi utils.PolicyChange
	var pol utils.Policy
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
	fmt.Println("CHANGE POLICY, ", actionApi)
	var isOkAction bool = true
	var game *utils.Game
	var ok bool
	var players int
	pol = utils.GetPolicy(actionApi.ID)
	var choosePol utils.PolicyValue
	for _, opt := range pol.PossibleValue2 {
		if opt.Value == actionApi.Value {
			choosePol = opt
		}
	}

	if game, ok = onGoingGames[actionApi.GameID]; ok {
		for players = range game.ListPlayers {
			if game.ListPlayers[players].PlayerID == actionApi.PlayerID {
				for _, playerPol := range game.ListPlayers[players].Policies {
					if playerPol.ActionName == actionApi.ID {
						if (playerPol.Value-actionApi.Value) > pol.MaxChange || (playerPol.Value-actionApi.Value) < -pol.MaxChange {
							fmt.Println("TOO MUICH CHANGE")
							isOkAction = false
						}
					}
				}
				if !utils.CheckConstraint(&game.ListPlayers[players], choosePol.Constraints, nil, game, 0) {
					fmt.Println("CONSTRAINT FAIL")

					isOkAction = false
				} else {
					fmt.Println("CONSTRAINT OK")
				}
			}
		}

	}
	if isOkAction {

		gMsg.Action = pol.ActionName
		gMsg.GameID = actionApi.GameID
		gMsg.PlayerID = actionApi.PlayerID
		gMsg.Text = "Order"
		gMsg.Effects = choosePol.Effects
		gMsg.Value = choosePol.Value
		gMsg.Type = "POLICY"
		jsonMsg, err := json.Marshal(gMsg)
		fmt.Println(string(jsonMsg))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(onGoingGames)
		ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}
	} else {
		fmt.Println("CANT CHANGE POLICY")
	}

}

func GetTranslations(w http.ResponseWriter, r *http.Request) {
	var translations []utils.Translation
	var language utils.Translation
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	fmt.Println(body)
	if err := json.Unmarshal(body, &language); err != nil {
		panic(err)
	}

	translations = utils.GetTranslationsByLanguage(language.Language)

	w.WriteHeader(http.StatusOK)
	jsonMsg, err := json.Marshal(translations)
	if err != nil {
		fmt.Println("fail :(")
		fmt.Println(err)
	}
	w.Write([]byte(jsonMsg))
}

func GetInfos(w http.ResponseWriter, r *http.Request) {
	var translations *[]utils.DisplayInfoCat

	translations = utils.GetInfos()

	w.WriteHeader(http.StatusOK)
	jsonMsg, err := json.Marshal(translations)
	if err != nil {
		fmt.Println("fail :(")
		fmt.Println(err)
	}
	w.Write([]byte(jsonMsg))
}

func GetHistory(w http.ResponseWriter, r *http.Request) {

	var acc utils.Account
	var accList []utils.Account
	var list []utils.GameHistory
	var apiList []utils.GameHistoryApi

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	fmt.Println(body)
	if err := json.Unmarshal(body, &acc); err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", ConnexionString)
	db.Where(&acc).First(&acc)
	db.Find(&accList)
	fmt.Println(acc.ID)
	db.Where("winner_id = ? OR loser_id = ?", acc.ID, acc.ID).Find(&list).Joins("JOIN accounts ON winner_id = accounts.ID OR loser_id = accounts.ID")
	rows, err := db.Table("game_histories").Select("game_histories.created_at, game_histories.elo_diff, winner.Name, loser.Name").Joins("JOIN accounts as winner ON winner_id = winner.ID").Joins("JOIN accounts as loser ON loser_id = loser.ID").Rows()
	fmt.Println("list", list, rows, err)
	for rows.Next() {
		var apiHist utils.GameHistoryApi
		rows.Scan(&apiHist.Created_at, &apiHist.ELODiff, &apiHist.WinnerNick, &apiHist.LoserNick)
		fmt.Println(apiHist)
		apiList = append(apiList, apiHist)
	}

	// rows, err := db.Table("game_histories").Select("game_histories.created_at, game_histories.elo_diff, winner.Name, loser.Name").Where("game_histories.winner_id = ? OR game_histories.loser_id = ?", acc.ID, acc.ID).Joins("JOIN accounts as winner ON winner_id = winner.ID OR winner_id = 0").Joins("JOIN accounts as loser ON loser_id = loser.ID OR loser_id = 0").Rows()
	// rows, err := db.Table("game_histories").Select("game_histories.created_at, game_histories.elo_diff, game_histories.winner_nick, game_histories.loser_id").Where("game_histories.winner_id = ? OR game_histories.loser_id = ?", acc.ID, acc.ID).Rows()
	fmt.Println("list", list, rows, err)
	for rows.Next() {
		var apiHist utils.GameHistoryApi
		rows.Scan(&apiHist.Created_at, &apiHist.ELODiff, &apiHist.WinnerNick, &apiHist.LoserNick)
		fmt.Println(apiHist)
		apiList = append(apiList, apiHist)
	}
	w.WriteHeader(http.StatusOK)
	jsonMsg, err := json.Marshal(apiList)
	if err != nil {
		fmt.Println("fail :(")
		fmt.Println(err)
	}
	w.Write([]byte(jsonMsg))
}

func GetLeaderBoard(w http.ResponseWriter, r *http.Request) {

	var accs []utils.Account
	var accsApi []utils.AccountLeaderboardApi
	db, err := gorm.Open("mysql", ConnexionString)
	db.Order("ELO desc, Name").Find(&accs)

	for _, i := range accs {
		accsApi = append(accsApi, utils.AccountLeaderboardApi{
			Name: i.Name,
			ELO:  i.ELO,
		})

	}
	w.WriteHeader(http.StatusOK)
	jsonMsg, err := json.Marshal(accsApi)
	if err != nil {
		fmt.Println("fail :(")
		fmt.Println(err)
	}
	w.Write([]byte(jsonMsg))
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
				if !utils.CheckConstraint(&game.ListPlayers[players], techno.Constraints, techno.Costs, game, 0) {
					fmt.Println("CONSTRAINT FAIL")
					isOkAction = false
				} else {
					fmt.Println("CONSTRAINT OK")
				}
			}
		}

	}
	if isOkAction {
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
				if !utils.CheckConstraint(&game.ListPlayers[players], action.Constraints, action.Costs, game, actionApi.Value) {
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
		if action.Selector == "range" {
			fmt.Println("RANGE ACTION", actionApi.Value)
			for i, e := range gMsg.Effects {
				fmt.Println("APPLY VALUE ON EFFECT")
				e.Value = actionApi.Value
				gMsg.Effects[i] = e

			}
		}

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
	db, err := gorm.Open("mysql", ConnexionString)
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

	fmt.Println("current games", onGoingGames)

	var isNewGame = true

	for _, g := range onGoingGames {
		for _, p := range g.ListPlayers {
			if p.Nick == acc.Name {
				isNewGame = false
			}
		}
	}

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

	if isNewGame {
		fmt.Println("NEW GAME")
		matchmakingQueue <- acc
	}

}

func JoinGameAi(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", ConnexionString)
	fmt.Println("Seems like someone want to join AI ! ", r.Body)
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

	fmt.Println("current games", onGoingGames)

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

	fmt.Println("NEW GAME")
	matchmakingAiQueue <- acc

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
