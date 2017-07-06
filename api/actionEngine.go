package main

import (
	"github.com/jinzhu/gorm"
	"github.com/orolol/utils"
)

func getDefaultPolicies() [][]utils.Policy {
	db, _ := gorm.Open("sqlite3", "test.db")
	var milPolicy []utils.Policy
	var ecoPolicy []utils.Policy
	db.Where("TypePolicy = ?", "MIL").Find(&milPolicy)
	db.Where("TypePolicy = ?", "ECO").Find(&ecoPolicy)

	return {milPolicy, ecoPolicy}

}
