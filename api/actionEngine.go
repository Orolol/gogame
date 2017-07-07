package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/orolol/gogame/utils"
)

func getDefaultPolicies() [][]utils.Policy {
	db, _ := gorm.Open("sqlite3", "test.db")
	var milPolicy []utils.Policy
	var ecoPolicy []utils.Policy
	db.Where("type_policy = ?", "MIL").Find(&milPolicy)
	db.Where("type_policy = ?", "ECO").Find(&ecoPolicy)

	fmt.Println("GET POLICIES")
	fmt.Println(milPolicy)
	fmt.Println(ecoPolicy)

	return [][]utils.Policy{milPolicy, ecoPolicy}

}
