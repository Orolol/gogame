package main

import (
	"fmt"

	"github.com/orolol/gogame/utils"
)

//InitializePlayerDefaultValue init player
func InitializePlayerDefaultValue(acc utils.Account) utils.PlayerInGame {
	army := utils.PlayerArmy{
		NbSoldier:  100000,
		NbLigtTank: 100,
		NbHvyTank:  50,
		NbArt:      50,
		NbAirSup:   0,
		NbAirBomb:  0,
		Morale:     100,
		Quality:    100}

	economy := utils.PlayerEconomy{
		Money:   100000000,
		TaxRate: 2}

	civilian := utils.PlayerCivilian{
		NbTotalCivil:       60000000,
		NbManpower:         600000,
		NbHeavyTankFactory: 20,
		NbLightTankFactory: 20,
		NbCivilianFactory:  20,
		NbResearchPoint:    0,
		NbScientist:        200}

	policy := utils.PlayerModifierPolicy{
		RecruitmentPolicy:  1,
		ManpowerSizePolicy: 1,
		ArtOnFactory:       false,
		BuildHvyTankFac:    true,
		BuildLgtTankFac:    true,
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

	var player = utils.PlayerInGame{
		PlayerID:       int(acc.ID),
		ModifierPolicy: policy,
		Army:           army,
		Nick:           acc.Name,
		Economy:        economy,
		Civilian:       civilian,
		Modifiers:      modifiers,
		Territory:      territory}

	return player
}

//PlayerAction player action
type PlayerAction func(player *utils.PlayerInGame, values map[string]float32)

//PASetRecruitementPolicy change recruitement policy to the value
func PASetRecruitementPolicy(player *utils.PlayerInGame, values map[string]float32) {
	qualityChange := player.ModifierPolicy.RecruitmentPolicy - values["value"]
	fmt.Println("QUALITY CHANGE ", qualityChange)
	player.Army.Quality -= values["value"]
	player.ModifierPolicy.RecruitmentPolicy = values["value"]
}
func setTaxRatePolicy(player *utils.PlayerInGame, values map[string]float32) {
	player.Economy.TaxRate = values["value"]
}
func setConscPolicy(player *utils.PlayerInGame, values map[string]float32) {
	player.Civilian.NbManpower -= player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
	player.Civilian.NbTotalCivil += player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
	player.ModifierPolicy.ManpowerSizePolicy = values["value"]
	player.Civilian.NbManpower += player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
	player.Civilian.NbTotalCivil -= player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
}
func setBuildLgtTank(player *utils.PlayerInGame, values map[string]float32) {
	if values["value"] == 1.0 {
		player.ModifierPolicy.BuildLgtTankFac = true
	} else {
		player.ModifierPolicy.BuildLgtTankFac = false
	}
}
func setBuildHvyTank(player *utils.PlayerInGame, values map[string]float32) {
	if values["value"] == 1.0 {
		player.ModifierPolicy.BuildHvyTankFac = true
	} else {
		player.ModifierPolicy.BuildHvyTankFac = false
	}
}

func actionCivConvertFactoryToLightTankFact(player *utils.PlayerInGame, values map[string]float32) {
	if player.Civilian.NbCivilianFactory > values["value"] {
		player.Civilian.NbCivilianFactory -= values["value"]
		player.Civilian.NbLightTankFactory += values["value"]
		var order = utils.PlayerLastOrders{
			OrderID:  "actionCivConvertFactoryToLightTankFact",
			Cooldown: int(values["CD"]),
		}
		player.LastOrders = append(player.LastOrders, order)
	}
}
func actionCivConvertFactoryToHvyTankFact(player *utils.PlayerInGame, values map[string]float32) {
	if player.Civilian.NbCivilianFactory > values["value"] {
		player.Civilian.NbCivilianFactory -= values["value"]
		player.Civilian.NbHeavyTankFactory += values["value"]
		var order = utils.PlayerLastOrders{
			OrderID:  "actionCivConvertFactoryToHvyTankFact",
			Cooldown: int(values["CD"]),
		}
		player.LastOrders = append(player.LastOrders, order)
	}
}

func actionWarPropaganda(player *utils.PlayerInGame, values map[string]float32) {
	player.Army.Morale += 15
	var order = utils.PlayerLastOrders{
		OrderID:  "actionWarPropaganda",
		Cooldown: int(values["CD"]),
	}
	player.LastOrders = append(player.LastOrders, order)

}
func emergencyRecruitment(player *utils.PlayerInGame, values map[string]float32) {
	player.Army.Morale -= 10
	player.Army.NbSoldier += player.Civilian.NbManpower * 0.1
	var order = utils.PlayerLastOrders{
		OrderID:  "emergencyRecruitment",
		Cooldown: int(values["CD"]),
	}
	player.LastOrders = append(player.LastOrders, order)

}
func purgeSoldier(player *utils.PlayerInGame, values map[string]float32) {
	player.Army.Morale += 15
	player.Modifiers["soldierQuality"] *= 1.15
	player.Army.NbSoldier *= 0.85
	var order = utils.PlayerLastOrders{
		OrderID:  "purgeSoldier",
		Cooldown: int(values["CD"]),
	}
	player.LastOrders = append(player.LastOrders, order)

}
func buyForeignTanks(player *utils.PlayerInGame, values map[string]float32) {
	player.Army.NbHvyTank += 50
	player.Army.NbLigtTank += 150
	var order = utils.PlayerLastOrders{
		OrderID:  "buyForeignTanks",
		Cooldown: int(values["CD"]),
	}
	player.LastOrders = append(player.LastOrders, order)

}

func genericApplyEffect(player *utils.PlayerInGame, effects []utils.Effect) {
	for _, e := range effects {
		utils.ApplyEffect(player, e)
	}
}

func genericApplyCosts(player *utils.PlayerInGame, costs []utils.Cost) {
	for _, c := range costs {
		utils.ApplyCost(player, c)
	}
}

// func technoIndusT1N1(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoIndusT1N1")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.CivilianProduction += 0.15
// }
// func technoIndusT1N2(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoIndusT1N2")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.TankProduction += 0.15
// }
// func technoIndusT1N3(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoIndusT1N3")
// 	fmt.Println("No airplane yet :()")
// }
// func technoIndusT2N1(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoIndusT2N1")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.CivilianProduction += 0.15
// }
// func technoIndusT2N2(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoIndusT2N2")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.TankProduction += 0.15
// }
// func technoIndusT2N3(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoIndusT2N3")
// 	fmt.Println("No airplane yet :()")
// }

// func technoMilT1N1(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoMilT1N1")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.CivilianProduction += 0.15
// }
// func technoMilT1N2(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoMilT1N2")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.TankProduction += 0.15
// }
// func technoMilT1N3(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoMilT1N3")
// 	fmt.Println("No airplane yet :()")
// }
// func technoMilT2N1(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoMilT2N1")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.CivilianProduction += 0.15
// }
// func technoMilT2N2(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoMilT2N2")
// 	player.Civilian.NbResearchPoint -= values["value"]
// 	player.ModifierPolicy.TankProduction += 0.15
// }
// func technoMilT2N3(player *utils.PlayerInGame, values map[string]float32) {
// 	player.PlayerTechnology = append(player.PlayerTechnology, "technoMilT2N3")
// 	fmt.Println("No airplane yet :()")
// }
