package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/orolol/utils"
)


/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc utils.Account
	fmt.Println("CREATE ACC")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.Unmarshal(r, &acc); err != nil {
		fmt.Println(acc)
		db.Create(&acc)
		fmt.Println("CREATED")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(acc.ID))
	} else {
		fmt.Println("FAIL CREATION")
		w.WriteHeader(http.StatusForbidden)
	}
}
