package main

import (
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
		TaxRate: 1}

	civilian := utils.PlayerCivilian{
		NbTotalCivil:       61000000,
		NbManpower:         600000,
		NbHeavyTankFactory: 20,
		NbLightTankFactory: 20,
		NbCivilianFactory:  20,
		NbResearchPoint:    0,
		NbScientist:        200}

	policy := utils.PlayerModifierPolicy{
		RecruitmentPolicy:  1,
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

	var player = utils.PlayerInGame{
		PlayerID:       int(acc.ID),
		ModifierPolicy: policy,
		Army:           army,
		Nick:           acc.Name,
		Economy:        economy,
		Civilian:       civilian,
		Modifiers:      modifiers,
		Territory:      territory,
		Policies:       policies}

	return player
}

//PlayerAction player action
type PlayerAction func(player *utils.PlayerInGame, values float32)

//setPopRecPolicy change recruitement policy to the value
func setPopRecPolicy(player *utils.PlayerInGame, values float32) {
	// qualityChange := player.ModifierPolicy.RecruitmentPolicy - values
	// player.Army.Quality -= values
	player.ModifierPolicy.RecruitmentPolicy = values
}
func setTaxRatePolicy(player *utils.PlayerInGame, values float32) {
	player.Economy.TaxRate = values
}
func setConscPolicy(player *utils.PlayerInGame, values float32) {
	player.Civilian.NbManpower -= player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
	player.Civilian.NbTotalCivil += player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
	player.ModifierPolicy.ManpowerSizePolicy = values
	player.Civilian.NbManpower += player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
	player.Civilian.NbTotalCivil -= player.Civilian.NbTotalCivil * player.ModifierPolicy.ManpowerSizePolicy * 0.01
}
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

func actionCivConvertFactoryToLightTankFact(player *utils.PlayerInGame, values float32) {
	if player.Civilian.NbCivilianFactory > values {
		player.Civilian.NbCivilianFactory -= values
		player.Civilian.NbLightTankFactory += values

	}
}
func actionCivConvertFactoryToHvyTankFact(player *utils.PlayerInGame, values float32) {
	if player.Civilian.NbCivilianFactory > values {
		player.Civilian.NbCivilianFactory -= values
		player.Civilian.NbHeavyTankFactory += values

	}
}

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
