package utils

import "github.com/jinzhu/gorm"

func SetBaseValueDB() {
	db, _ := gorm.Open("sqlite3", "test.db")
	db.DropTable(&Policy{})
	db.CreateTable(&Policy{})

	var popRecPol Policy = Policy{
		Name:           "Recruitement Policy",
		ActionName:     "setPopRecPolicy",
		ConstraintName: "consPopRecPolicy",
		Description:    "Set your recuitement policy",
		TypePolicy:     "MIL",
		PossibleValue:  "['0,01','0,02','0,05','0,07','0,1']",
		DefaultValue:   "0,01"}
	db.Create(&popRecPol)

	var taxRatePol Policy = Policy{
		Name:           "Tax rate",
		ActionName:     "setTaxRatePolicy",
		ConstraintName: "consTaxRatePolicy",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "['0,01','0,02','0,05','0,07','0,1']",
		DefaultValue:   "0,05"}
	db.Create(&taxRatePol)

}
