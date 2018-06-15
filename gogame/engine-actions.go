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
		Barracks:     50,
		SmallCities:  100,
		MediumCities: 25,
		BigCities:    5,
		Surface:      100,
	}

	var modifiers = make(map[string]float32)

	//MIL modifiers
	modifiers["soldierRecruitmentExperience"] = 1.0
	modifiers["workersConcrptionEfficiency"] = 1.0

	modifiers["researchEfficiency"] = 1.0

	modifiers["soldierQuality"] = 1.0
	modifiers["lightTankQuality"] = 1.0
	modifiers["heavyTankQuality"] = 1.0

	modifiers["soldierArmor"] = 1.0
	modifiers["lightTankArmor"] = 1.0
	modifiers["heavyTankArmor"] = 1.0

	//PROD modifiers
	modifiers["civilianFactoryProduction"] = 1.0
	modifiers["lightTankFactoryProduction"] = 1.0
	modifiers["heavyTankFactoryProduction"] = 1.0

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

func actionWarPropaganda(player *utils.PlayerInGame, values float32) {
	player.Army.Morale += 15

}
func emergencyRecruitment(player *utils.PlayerInGame, values float32) {
	player.Army.Morale -= 10
	player.Army.NbSoldier += player.Civilian.NbManpower * 0.1

}
func purgeSoldier(player *utils.PlayerInGame, values float32) {
	player.Army.Morale += 15
	player.Modifiers["soldierQuality"] *= 1.15
	player.Army.NbSoldier *= 0.85

}
func buyForeignTanks(player *utils.PlayerInGame, values float32) {
	player.Army.NbHvyTank += 50
	player.Army.NbLigtTank += 150

}

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
