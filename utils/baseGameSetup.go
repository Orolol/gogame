package utils

import "github.com/jinzhu/gorm"

func SetBaseValueDB() {
	db, _ := gorm.Open("sqlite3", "test.db")
	db.DropTable(&Policy{})
	db.CreateTable(&Policy{})
	db.DropTable(&PlayerActionOrder{})
	db.CreateTable(&PlayerActionOrder{})

	var popRecPol Policy = Policy{
		Name:           "Training Time",
		ActionName:     "setPopRecPolicy",
		ConstraintName: "consPopRecPolicy",
		Description:    "Set your recuitement policy",
		TypePolicy:     "MIL",
		PossibleValue:  "{\"Full\" : 1,\"Long\" : 2,\"Hurry\" : 5,\"No time !\" : 10,\"Send everyone !\" : 30}",
		DefaultValue:   "1"}
	db.Create(&popRecPol)
	var conscPol Policy = Policy{
		Name:           "Conscription Policy",
		ActionName:     "setConscPolicy",
		ConstraintName: "consPopRecPolicy",
		Description:    "Set your recuitement policy",
		TypePolicy:     "MIL",
		PossibleValue:  "{\"Pro Army\" : 1,\"Volonteer\" : 2,\"War time\" : 5,\"All valids !\" : 10,\"Anyone who can hold a weapon\" : 30}",
		DefaultValue:   "1"}
	db.Create(&conscPol)

	var taxRatePol Policy = Policy{
		Name:           "Tax rate",
		ActionName:     "setTaxRatePolicy",
		ConstraintName: "consTaxRatePolicy",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "{\"Low taxes\" : 1,\"Country effort\" : 1.5,\"War Economy\" : 2,\"Full Mobilization\" : 3,\"Total war\" : 5}",
		DefaultValue:   "1.5"}
	db.Create(&taxRatePol)

	var lgtTankBuild Policy = Policy{
		Name:           "Build Light Tank ?",
		ActionName:     "setBuildLgtTank",
		ConstraintName: "consTaxRatePolicy",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "{\"Yes\" : 1,\"No\" : 0}",
		DefaultValue:   "1"}
	db.Create(&lgtTankBuild)

	var hvyTankBuild Policy = Policy{
		Name:           "Build Heavy Tank ?",
		ActionName:     "setBuildHvyTank",
		ConstraintName: "consTaxRatePolicy",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "{\"Yes\" : 1,\"No\" : 0}",
		DefaultValue:   "1"}
	db.Create(&hvyTankBuild)

	var civToLight PlayerActionOrder = PlayerActionOrder{
		Name:           "Civ Fact -> Lght Fact",
		ActionName:     "actionCivConvertFactoryToLightTankFact",
		ConstraintName: "actionCivConvertFactoryToLightTankFact",
		Description:    "Convert Civilian Factory to light Tank factory (Cost 1M) ",
	}
	db.Create(&civToLight)

	var civToHvy PlayerActionOrder = PlayerActionOrder{
		Name:           "Civ Fact -> Hvy Fact",
		ActionName:     "actionCivConvertFactoryToHvyTankFact",
		ConstraintName: "actionCivConvertFactoryToHvyTankFact",
		Description:    "Convert Civilian Factory to Heavy Tank factory (Cost 1M) ",
	}
	db.Create(&civToHvy)
	var warProp PlayerActionOrder = PlayerActionOrder{
		Name:           "War Propaganda",
		ActionName:     "actionWarPropaganda",
		ConstraintName: "actionWarPropaganda",
		Description:    "Boost morale by 15% (cost 10M) ",
	}
	db.Create(&warProp)

}
