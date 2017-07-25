package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/orolol/gogame/utils"
)

func getDefaultPolicies() map[string][]utils.Policy {
	db, _ := gorm.Open("sqlite3", "test.db")
	var milPolicy []utils.Policy
	var ecoPolicy []utils.Policy
	db.Where("type_policy = ?", "MIL").Find(&milPolicy)
	db.Where("type_policy = ?", "ECO").Find(&ecoPolicy)

	var ret = make(map[string][]utils.Policy)
	fmt.Println(milPolicy)
	fmt.Println(ecoPolicy)
	ret["MIL"] = milPolicy
	ret["ECO"] = ecoPolicy

	fmt.Println(ret)

	return ret

}

func getDefaultActions() []utils.PlayerActionOrder {
	db, _ := gorm.Open("sqlite3", "test.db")
	var actions []utils.PlayerActionOrder
	db.Find(&actions)
	return actions
}

func getDefaultTech() []utils.Technology {
	db, _ := gorm.Open("sqlite3", "test.db")
	var actions []utils.Technology
	db.Find(&actions)
	return actions
}
