package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/orolol/gogame/utils"
)

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://http://localhost:8080/todos

*/
func SignUp(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", ConnexionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var acc utils.Account
	fmt.Println("CREATE ACC")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if erra := r.Body.Close(); erra != nil {
		panic(erra)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if errb := json.Unmarshal(body, &acc); errb != nil {
		fmt.Println("FAIL CREATION ", errb)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	fmt.Println(acc)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("PASS : ", acc.Password)
	acc.Password = string(hashedPassword)
	fmt.Println("FINAL PASS : ", acc.Password)

	acc.ELO = 1500

	stoken, err := utils.GenerateRandomString(22)
	if err != nil {
		fmt.Println("Error while generating token")
		return
	}
	token := utils.Token{Token: stoken, Status: "active"}

	if err := db.Create(&token).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	acc.TokenID = token.ID

	if err := db.Create(&acc).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println("CREATED ", acc.ID)
		fmt.Println("CREATED ", acc.Name)
		fmt.Println("CREATED ", acc.Login)
		fmt.Println("CREATED ", acc.Password)
		fmt.Println("CREATED ", acc.TokenID)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(stoken))
	}

}

func EditAccount(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", ConnexionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var acc utils.Account
	var dbacc utils.Account
	fmt.Println("EDIT ACC")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if erra := r.Body.Close(); erra != nil {
		panic(erra)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if errb := json.Unmarshal(body, &acc); errb != nil {
		fmt.Println("FAIL CREATION ", errb)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	fmt.Println(acc)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("PASS : ", acc.Password)
	acc.Password = string(hashedPassword)
	fmt.Println("FINAL PASS : ", acc.Password)

	if res := db.First(&dbacc, "Login = ?", acc.Login); res.Error != nil {
		fmt.Println("ERROR EDIT", err.Error)
		fmt.Println("ERROR EDIT", acc.Login)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		fmt.Println("EDIT THIS", dbacc)
		dbacc.Name = acc.Name
		if acc.Password != "" {
			dbacc.Password = acc.Password
		}

		db.Save(&dbacc)
		w.WriteHeader(http.StatusCreated)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK LETS LOGIN")
	db, err := gorm.Open("mysql", ConnexionString)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var acc utils.Account
	var accApi utils.AccountApi

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &acc); err != nil {
		fmt.Println("erf")
	}

	clearPass := acc.Password

	db.First(&acc, "Login = ?", acc.Login)
	errPass := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(clearPass))

	if errPass != nil {
		fmt.Println("Mauvais password", errPass, acc.Password, clearPass)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bad password"))
	} else if acc.ID == 0 {
		fmt.Println("Mauvais account")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Bad account"))
	} else {
		w.WriteHeader(http.StatusOK)

		var token utils.Token

		fmt.Println("LOGGED ", acc.ID)

		db.Model(&acc).Related(&token)

		accApi.ID = acc.ID
		accApi.Login = acc.Login
		accApi.Name = acc.Name
		accApi.ELO = acc.ELO
		accApi.Token = token.Token

		jsonMsg, err := json.Marshal(accApi)
		if err != nil {
			fmt.Println("fail :(")
			fmt.Println(err)
		}
		fmt.Println("ACC ", accApi, jsonMsg)
		w.Write([]byte(jsonMsg))
	}
}
