package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

//CheckConstraint Check if constraint is respected.
func CheckConstraint(player *PlayerInGame, constraint string) bool {
	fmt.Println("CONSTRAINT CHECK ", constraint, player.Nick)
	var conObj Constraint
	err := json.Unmarshal([]byte(constraint), &conObj)
	if err == nil {
		for _, t := range conObj.Tech {
			if !stringInSlice(t, player.PlayerTechnology) {
				return false
			}
		}
	} else {
		return false
	}
	return true

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
	
//ApplyEffect apply effect on player modifiers
func ApplyEffect(player *PlayerInGame, effect Effect) {
	for _, mod := range player.Modifiers {
		if mod.Name == effect.ModifierName {
			switch op := effect.Operator; op {
			case "+":
				mod.Value += effect.Value
			case "-":
				mod.Value -= effect.Value
			case "*":
				mod.Value *= effect.Value
			case "/":
				mod.Value *= 1 / effect.Value

			}
		}
	}
}

//AlgoDamageDealt Calculate dmg dealt
func AlgoDamageDealt(player *PlayerInGame) float32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rollp1 = r1.Float32() + 4.0
	var rollp2 = r1.Float32() + 4.0
	var rollp3 = r1.Float32() + 4.0
	dmgModifier := (player.Army.Morale / 100.0) * (player.Army.Quality / 100.0)
	var dmg = ((player.Army.NbSoldier * 0.05 * rollp1) + (player.Army.NbLigtTank * 5 * rollp2) + (player.Army.NbHvyTank * 15 * rollp3)) * 0.2 * dmgModifier
	fmt.Println("DAMAGE ", dmg)
	return dmg
}

//AlgoReinforcement calc reinforcement
func AlgoReinforcement(player *PlayerInGame) *PlayerInGame {
	if player.Economy.Money > 0 {
		var minRf float32 = 100.0
		reinforcement := 1000 * player.ModifierPolicy.RecruitmentPolicy
		fmt.Println("REINFORCEMENT : ", reinforcement)
		if reinforcement > player.Civilian.NbManpower {
			reinforcement = 0.0
		} else if reinforcement < minRf {
			reinforcement = minRf
		}

		player.Army.NbSoldier += reinforcement
		player.Civilian.NbManpower -= reinforcement
	}
	natGrowth := player.ModifierPolicy.ManpowerSizePolicy * 0.00001 * player.Civilian.NbTotalCivil
	player.Civilian.NbManpower += natGrowth
	player.Civilian.NbTotalCivil -= natGrowth

	fmt.Println("NAT GROTH ", natGrowth)

	return player
}

//AlgoDamageRepartition Calculate loses
func AlgoDamageRepartition(player *PlayerInGame, dmgIncoming float32) *PlayerInGame {
	totalHp := player.Army.NbSoldier + (player.Army.NbLigtTank * 5) + (player.Army.NbHvyTank * 20)
	var multiHvyTank float32 = 0.0
	var multiLgtTank float32 = 0.0
	if player.Army.NbHvyTank > 0 {

		multiHvyTank = (player.Army.NbHvyTank * 20) / totalHp
	}
	if player.Army.NbLigtTank > 0 {
		multiLgtTank = (player.Army.NbLigtTank * 5) / totalHp
	}
	multiSoldier := 1 - multiHvyTank - multiLgtTank

	if dmgIncoming > totalHp {
		fmt.Println("Civilian damage ", dmgIncoming-totalHp)
	}
	fmt.Println("DMG MODIFER", (player.Army.Morale/100.0)*(player.Army.Quality/100.0))
	dmgModifier := 2 / (1 + (player.Army.Morale/100.0)*(player.Army.Quality/100.0))
	player.Army.NbSoldier -= dmgIncoming * multiSoldier * 0.1 * dmgModifier
	player.Army.NbLigtTank -= dmgIncoming * multiLgtTank * 0.02 * dmgModifier
	player.Army.NbHvyTank -= dmgIncoming * multiHvyTank * 0.005 * dmgModifier

	if player.Army.NbSoldier < 0 {
		player.Army.NbSoldier = 0.0
	}
	if player.Army.NbLigtTank < 0 {
		fmt.Println("All lght tank lost ")
		player.Army.NbLigtTank = 0.0
	}
	if player.Army.NbHvyTank < 0 {
		fmt.Println("All hvy tank lost ")
		player.Army.NbHvyTank = 0.0
	}

	return player
}

func AlgoEconomicEndTurn(player *PlayerInGame) *PlayerInGame {
	armyUpkeep := (player.Army.NbSoldier * 100) + (player.Army.NbLigtTank * 150) + (player.Army.NbHvyTank * 200)
	tax := (player.Economy.TaxRate * 0.2 * player.Civilian.NbTotalCivil)
	fmt.Println("MONEY : ", player.Economy.Money)
	fmt.Println("armyUpkeep : ", armyUpkeep)
	fmt.Println("tax : ", tax)
	player.Economy.Money = player.Economy.Money - armyUpkeep + tax

	//Technology

	player.Civilian.NbResearchPoint += player.Civilian.NbScientist * 0.05

	if player.Economy.Money > 0 {

		player.Army.NbLigtTank = player.Army.NbLigtTank + player.Civilian.NbLightTankFactory*3
		player.Army.NbHvyTank = player.Army.NbHvyTank + player.Civilian.NbHeavyTankFactory*1

		player.Economy.Money -= (player.Civilian.NbLightTankFactory * 10000) + (player.Civilian.NbHeavyTankFactory * 100000)

		var nbThingToBuild float32 = 1.0

		var civilianProduction = player.Civilian.NbCivilianFactory * 0.01 * (2 / player.Economy.TaxRate) * (2 / player.ModifierPolicy.ManpowerSizePolicy)
		fmt.Println("CIVILIAN PROD :", civilianProduction)
		if player.ModifierPolicy.BuildLgtTankFac {
			nbThingToBuild += 1.0
		}
		if player.ModifierPolicy.BuildHvyTankFac {
			nbThingToBuild += 1
		}
		if player.ModifierPolicy.BuildLgtTankFac {
			player.Civilian.NbLightTankFactory += (civilianProduction / nbThingToBuild) * 0.5
		}
		if player.ModifierPolicy.BuildHvyTankFac {
			player.Civilian.NbHeavyTankFactory += (civilianProduction / nbThingToBuild) * 0.4
		}
		player.Civilian.NbCivilianFactory += (civilianProduction / nbThingToBuild) * 0.2

	} else {
		if player.Army.Morale > 10 {
			player.Army.Morale -= 1
		}

		fmt.Println("ENOUGHT MONEY TO BUILD !")
		fmt.Println("MORALE !", player.Army.Morale)
	}

	return player
}
