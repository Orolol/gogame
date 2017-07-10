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
	db, _ := gorm.Open("sqlite3", "test.db")

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

	db.Where("ID = ?", polChange.ID).First(&pol)
	gMsg.Action = pol.ActionName
	gMsg.GameID = polChange.GameID
	gMsg.PlayerID = polChange.PlayerID
	gMsg.Text = "Change pol"

	gMsg.Value = polChange.Value

	jsonMsg, err := json.Marshal(gMsg)
	fmt.Println(string(jsonMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SEND Game MSG!")
	ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}

}

func Actions(w http.ResponseWriter, r *http.Request) {
	var actionApi utils.PolicyChange
	var action utils.PlayerActionOrder
	var gMsg utils.GameMsg
	db, _ := gorm.Open("sqlite3", "test.db")

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

	db.Where("ID = ?", actionApi.ID).First(&action)
	gMsg.Action = action.ActionName
	gMsg.GameID = actionApi.GameID
	gMsg.PlayerID = actionApi.PlayerID
	gMsg.Text = "Change pol"

	gMsg.Value = actionApi.Value

	jsonMsg, err := json.Marshal(gMsg)
	fmt.Println(string(jsonMsg))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("SEND Game MSG! ", gMsg)
	ZMQPusher.SendChan <- [][]byte{[]byte("MSG"), []byte(jsonMsg)}

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
