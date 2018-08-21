package main

import (
	"github.com/orolol/gogame/utils"
)

//InitializePlayerDefaultValue init player
func InitializePlayerDefaultValue(acc utils.Account) utils.PlayerInGame {
	army := utils.PlayerArmy{
		NbSoldier:         30000,
		NbLigtTank:        200,
		NbHvyTank:         100,
		NbArt:             50,
		NbAirSup:          0,
		NbAirBomb:         0,
		Ammunition:        100000,
		InfantryEquipment: 100000,
		Morale:            100,
		Quality:           100}

	economy := utils.PlayerEconomy{
		Money:                       10000000,
		TaxRate:                     1,
		LightTankProduction:         100,
		InfantryEquipmentProduction: 100,
		HeavyTankProduction:         100,
		FighterProduction:           100,
		ArtilleryProduction:         100,
		BomberProduction:            100,
		FactoryProduction:           100,
		AmmunitionProduction:        100,
	}

	civilian := utils.PlayerCivilian{
		NbManpower:        300000,
		NbCivilianFactory: 30,
		NbResearchPoint:   0,
		NbScientist:       200}

	policy := utils.PlayerModifierPolicy{
		TrainingPolicy:     1,
		ManpowerSizePolicy: 1,
		ArtOnFactory:       1,
		BuildHvyTankFac:    1,
		BuildLgtTankFac:    1,
		AirCraftProduction: 1,
		CivilianProduction: 1,
		TankProduction:     1}

	territory := utils.PlayerTerritory{
		Barracks:     45,
		SmallCities:  100,
		MediumCities: 22,
		BigCities:    5,
		Surface:      100,
	}

	var modifiers = make(map[string]float32)

	//MIL modifiers
	modifiers["soldierRecruitmentExperience"] = 1.0
	modifiers["workersConcrptionEfficiency"] = 1.0

	modifiers["researchEfficiency"] = 1.0

	modifiers["soldierQuality"] = 100
	modifiers["lightTankQuality"] = 100
	modifiers["heavyTankQuality"] = 100

	modifiers["soldierArmor"] = 1.0
	modifiers["lightTankArmor"] = 1.0
	modifiers["heavyTankArmor"] = 1.0

	//PROD modifiers
	modifiers["civilianFactoryProduction"] = 1.0
	modifiers["lightTankFactoryProduction"] = 1.0
	modifiers["heavyTankFactoryProduction"] = 1.0

	modifiers["bomberTargetArmy"] = 1.0
	modifiers["bomberTargetFactories"] = 0
	modifiers["bomberTargetPopulation"] = 0

	modifiers["engageFighter"] = 0
	modifiers["engageBomber"] = 0
	modifiers["engageAerialForce"] = 0

	modifiers["dmgAerialBonus"] = 100
	modifiers["dmgBombBonus"] = 100

	var policies []utils.PolicyValue
	//fmt.Println("GET POLICIES")
	for _, p := range utils.GetPolicies() {
		//fmt.Println("P", p)
		for _, pv := range p.PossibleValue2 {
			//fmt.Println("PV", pv)
			if pv.IsDefault {
				//fmt.Println("DEFAULT", pv)
				policies = append(policies, pv)
			}
		}
	}

	var infos = make(map[string]*utils.PlayerInformation)

	infos["civilianProduction"] = &utils.PlayerInformation{
		Type:        "INDUSTRIAL",
		SubType:     "FACTORY",
		Name:        "civilianProduction",
		Description: "Civilian production",
	}
	infos["TaxRate"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "TaxRate",
		Description: "Tax rate",
	}
	infos["TaxRevenu"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "TaxRevenu",
		Description: "Tax revenu",
	}

	infos["lightTankProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "lightTankProduction",
		Description: "Production of light tank each turn",
	}
	infos["heavyTankProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "heavyTankProduction",
		Description: "Production of heavy tank each turn",
	}
	infos["bomberProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "bomberProduction",
		Description: "Production of bomber each turn",
	}
	infos["fighterProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "fighterProduction",
		Description: "Production of fighter each turn",
	}
	infos["artilleryProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "artilleryProduction",
		Description: "Production of artillery each turn",
	}
	infos["factoryProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "factoryProduction",
		Description: "Production of factory each turn",
	}
	infos["infantryEquipmentProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "infantryEquipmentProduction",
		Description: "Production of infantry equipment each turn",
	}
	infos["ammunitionProduction"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "ammunitionProduction",
		Description: "Production of ammunition each turn",
	}
	infos["Factory"] = &utils.PlayerInformation{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Factory",
		Description: "Number of factories",
	}

	infos["SoldierDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Soldier damage dealt",
		Description: "Soldier damage dealt",
	}

	infos["LightTankDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Light Tank damage dealt",
		Description: "Light Tank damage dealt",
	}

	infos["HeavyTankDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Heavy Tank damage dealt",
		Description: "Heavy Tank damage dealt",
	}

	infos["ArtDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Artillery damage dealt",
		Description: "Artillery damage dealt",
	}

	infos["AirSupAerialDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Figther aerial damage dealt",
		Description: "Figther aerial damage dealt",
	}
	infos["AirSupGroundDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Figther ground damage dealt",
		Description: "Figther ground damage dealt",
	}
	infos["AirBombAerialDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Bomber aerial damage dealt",
		Description: "Bomber aerial damage dealt",
	}
	infos["AirBombGroundDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Bomber ground damage dealt",
		Description: "Bomber ground damage dealt",
	}

	infos["TotalAerialDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Total aerial damage dealt",
		Description: "Total aerial damage dealt",
	}
	infos["TotalGroundDmg"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Total ground damage dealt",
		Description: "Total ground damage dealt",
	}

	infos["SoldierRecruit"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "RECRUITMENT",
		Name:        "Total soldier recruited",
		Description: "Total soldier recruited",
	}
	infos["ManpowerGrowth"] = &utils.PlayerInformation{
		Type:        "MILITARY",
		SubType:     "RECRUITMENT",
		Name:        "Manpower Growth",
		Description: "Manpower Growth",
	}
	var cb = utils.GetCountry(acc.SelectedCountry)

	var player = utils.PlayerInGame{
		PlayerID:           int(acc.ID),
		ModifierPolicy:     policy,
		Army:               army,
		Nick:               acc.Name,
		Economy:            economy,
		Civilian:           civilian,
		Modifiers:          modifiers,
		Territory:          territory,
		Policies:           policies,
		PlayerInformations: infos,
		Country:            cb,
	}

	return player
}

//PlayerAction player action
type PlayerAction func(player *utils.PlayerInGame, values float32)

//setPopRecPolicy change recruitement policy to the value
func setPopRecPolicy(player *utils.PlayerInGame, values float32) {
	// qualityChange := player.ModifierPolicy.TrainingPolicy - values
	// player.Army.Quality -= values
	player.ModifierPolicy.TrainingPolicy = values
}
func setTaxRatePolicy(player *utils.PlayerInGame, values float32) {
	player.Economy.TaxRate = values
}

//
func setBuildLgtTank(player *utils.PlayerInGame, values float32) {
	if values == 1.0 {
		player.ModifierPolicy.BuildLgtTankFac = 1
	} else {
		player.ModifierPolicy.BuildLgtTankFac = 0
	}
}
func setBuildHvyTank(player *utils.PlayerInGame, values float32) {
	if values == 1.0 {
		player.ModifierPolicy.BuildHvyTankFac = 1
	} else {
		player.ModifierPolicy.BuildHvyTankFac = 0
	}
}

// func actionCivConvertFactoryToLightTankFact(player *utils.PlayerInGame, values float32) {
// 	if player.Civilian.NbCivilianFactory > values {
// 		player.Civilian.NbCivilianFactory -= values
// 		player.Civilian.NbLightTankFactory += values

// 	}
// }
// func actionCivConvertFactoryToHvyTankFact(player *utils.PlayerInGame, values float32) {
// 	if player.Civilian.NbCivilianFactory > values {
// 		player.Civilian.NbCivilianFactory -= values
// 		player.Civilian.NbHeavyTankFactory += values

// 	}
// }

func genericApplyEffect(player *utils.PlayerInGame, opponent *utils.PlayerInGame, effects []utils.Effect, game *utils.Game) {
	for _, e := range effects {
		if e.Target == "Player" || e.Target == "Both" || e.Target == "" {
			utils.ApplyEffect(player, e, game)
		}
		if e.Target == "Opponent" || e.Target == "Both" {
			utils.ApplyEffect(opponent, e, game)
		}
	}
}

func genericApplyCosts(player *utils.PlayerInGame, costs []utils.Cost) {
	for _, c := range costs {
		utils.ApplyCost(player, c)
	}
}
