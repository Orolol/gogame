package utils

import "github.com/jinzhu/gorm"

func SetBaseValueDB() {
	db, _ := gorm.Open("sqlite3", "test.db")
	db.DropTable(&Policy{})
	db.CreateTable(&Policy{})
	db.DropTable(&PlayerActionOrder{})
	db.CreateTable(&PlayerActionOrder{})
	db.DropTable(&Technology{})
	db.CreateTable(&Technology{})

	//POLICIES
	var popRecPol = Policy{
		Name:           "Training Time",
		ActionName:     "setPopRecPolicy",
		ConstraintName: "{}",
		Description:    "Set your recuitement policy",
		TypePolicy:     "MIL",
		PossibleValue:  "{\"Full\" : 1,\"Long\" : 2,\"Hurry\" : 5,\"No time !\" : 10,\"Send everyone !\" : 30}",
		DefaultValue:   "1"}
	db.Create(&popRecPol)
	var conscPol = Policy{
		Name:           "Conscription Policy",
		ActionName:     "setConscPolicy",
		ConstraintName: "{}",
		Description:    "Set your recuitement policy",
		TypePolicy:     "MIL",
		PossibleValue:  "{\"Pro Army\" : 1,\"Volonteer\" : 2,\"War time\" : 5,\"All valids !\" : 10,\"Anyone who can hold a weapon\" : 30}",
		DefaultValue:   "1"}
	db.Create(&conscPol)

	var taxRatePol = Policy{
		Name:           "Tax rate",
		ActionName:     "setTaxRatePolicy",
		ConstraintName: "{}",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "{\"Low taxes\" : 1,\"Country effort\" : 1.5,\"War Economy\" : 2,\"Full Mobilization\" : 3,\"Total war\" : 5}",
		DefaultValue:   "1.5"}
	db.Create(&taxRatePol)

	var lgtTankBuild = Policy{
		Name:           "Build Light Tank ?",
		ActionName:     "setBuildLgtTank",
		ConstraintName: "{}",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "{\"Yes\" : 1,\"No\" : 0}",
		DefaultValue:   "1"}
	db.Create(&lgtTankBuild)

	var hvyTankBuild = Policy{
		Name:           "Build Heavy Tank ?",
		ActionName:     "setBuildHvyTank",
		ConstraintName: "{}",
		Description:    "Set your tax rate. ",
		TypePolicy:     "ECO",
		PossibleValue:  "{\"Yes\" : 1,\"No\" : 0}",
		DefaultValue:   "1"}
	db.Create(&hvyTankBuild)

	//ACTIONS
	var CivToLight = PlayerActionOrder{
		Name:           "Civ Fact -> Lght Fact",
		ActionName:     "actionCivConvertFactoryToLightTankFact",
		ConstraintName: "{}",
		Description:    "Convert Civilian Factory to light Tank factory (Cost 1M) ",
		Cooldown:       10,
		Cost:           1000000,
	}
	db.Create(&CivToLight)

	var CivToHvy = PlayerActionOrder{
		Name:           "Civ Fact -> Hvy Fact",
		ActionName:     "actionCivConvertFactoryToHvyTankFact",
		ConstraintName: "{}",
		Description:    "Convert Civilian Factory to Heavy Tank factory (Cost 1M) ",
		Cooldown:       10,
		Cost:           1000000,
	}
	db.Create(&CivToHvy)
	var WarProp = PlayerActionOrder{
		Name:           "War Propaganda",
		ActionName:     "actionWarPropaganda",
		ConstraintName: "{}",
		Description:    "Boost morale by 15% (cost 10M) ",
		Cooldown:       10,
		Cost:           10000000,
	}
	db.Create(&WarProp)
	var buyFT = PlayerActionOrder{
		Name:           "Buy foreign tanks",
		ActionName:     "buyForeignTanks",
		ConstraintName: "{\"tech\":[\"technoIndusT1N1\", \"technoIndusT1N2\"]}",
		Description:    "Get 150 light tank and 50 heavy one (cost 100M) ",
		Cooldown:       60,
		Cost:           100000000,
	}
	db.Create(&buyFT)

	//TECHNOLOGY
	var technoIndusT1N1 = Technology{
		Name:           "Boost Civilian production",
		Description:    "Boost civilian factory production by 15%",
		Cost:           150.0,
		ActionName:     "technoIndusT1N1",
		ConstraintName: "{}",
		Tier:           1,
		TypeTechnology: "INDUS",
	}
	db.Create(&technoIndusT1N1)
	var technoIndusT1N2 = Technology{
		Name:           "Boost Tank production",
		Description:    "Boost Tank factory production by 15%",
		Cost:           100.0,
		ActionName:     "technoIndusT1N2",
		ConstraintName: "{}",
		Tier:           1,
		TypeTechnology: "INDUS",
	}
	db.Create(&technoIndusT1N2)
	var technoIndusT1N3 = Technology{
		Name:           "Boost Aircraft production",
		Description:    "Boost Aircraft factory production by 15%",
		Cost:           100.0,
		ActionName:     "technoIndusT1N3",
		ConstraintName: "{}",
		Tier:           1,
		TypeTechnology: "INDUS",
	}
	db.Create(&technoIndusT1N3)
	var technoIndusT2N1 = Technology{
		Name:           "Boost Civilian production T2",
		Description:    "Boost civilian factory production by 15%",
		Cost:           450.0,
		ActionName:     "technoIndusT2N1",
		ConstraintName: "{\"tech\":[\"technoIndusT1N1\"]}",
		Tier:           2,
		TypeTechnology: "INDUS",
	}
	db.Create(&technoIndusT2N1)
	var technoIndusT2N2 = Technology{
		Name:           "Boost Tank production T2",
		Description:    "Boost Tank factory production by 15%",
		Cost:           300.0,
		ActionName:     "technoIndusT2N2",
		ConstraintName: "{\"tech\":[\"technoIndusT1N2\"]}",
		Tier:           2,
		TypeTechnology: "INDUS",
	}
	db.Create(&technoIndusT2N2)
	var technoIndusT2N3 = Technology{
		Name:           "Boost Aircraft production T2",
		Description:    "Boost Aircraft factory production by 15%",
		Cost:           600.0,
		ActionName:     "technoIndusT2N3",
		ConstraintName: "{\"tech\":[\"technoIndusT1N3\"]}",
		Tier:           2,
		TypeTechnology: "INDUS",
	}
	db.Create(&technoIndusT2N3)

	//MIL TECHNOLOGY
	var technoMilT1N1 = Technology{
		Name:           "Boost soldier damage",
		Description:    "Boost soldier damage by 15%",
		Cost:           200.0,
		ActionName:     "technoMilT1N1",
		ConstraintName: "{}",
		Tier:           1,
		TypeTechnology: "MIL",
	}
	db.Create(&technoMilT1N1)

	var technoMilT1N2 = Technology{
		Name:           "Boost Tank damage",
		Description:    "Boost Tank damage by 15%",
		Cost:           100.0,
		ActionName:     "technoMilT1N2",
		ConstraintName: "{}",
		Tier:           1,
		TypeTechnology: "MIL",
	}
	db.Create(&technoMilT1N2)
	var technoMilT1N3 = Technology{
		Name:           "Boost Aircraft damage",
		Description:    "Boost Aircraft damage by 15%",
		Cost:           100.0,
		ActionName:     "technoMilT1N3",
		ConstraintName: "{}",
		Tier:           1,
		TypeTechnology: "MIL",
	}
	db.Create(&technoMilT1N3)
	var technoMilT2N1 = Technology{
		Name:           "Boost soldier damage T2",
		Description:    "Boost soldier damage by 15%",
		Cost:           800.0,
		ActionName:     "technoMilT2N1",
		ConstraintName: "{\"tech\":[\"technoMilT1N1\"]}",
		Tier:           2,
		TypeTechnology: "MIL",
	}
	db.Create(&technoMilT2N1)
	var technoMilT2N2 = Technology{
		Name:           "Boost Tank damage T2",
		Description:    "Boost Tank damage by 15%",
		Cost:           300.0,
		ActionName:     "technoMilT2N2",
		ConstraintName: "{\"tech\":[\"technoMilT1N2\"]}",
		Tier:           2,
		TypeTechnology: "MIL",
	}
	db.Create(&technoMilT2N2)
	var technoMilT2N3 = Technology{
		Name:           "Boost Aircraft damage T2",
		Description:    "Boost Aircraft damage by 15%",
		Cost:           600.0,
		ActionName:     "technoMilT2N3",
		ConstraintName: "{\"tech\":[\"technoMilT1N3\"]}",
		Tier:           2,
		TypeTechnology: "MIL",
	}
	db.Create(&technoMilT2N3)

}
