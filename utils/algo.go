package utils

import (
	"fmt"
	"math/rand"
	"time"
)

//CheckConstraint Check if constraint is respected.
func CheckConstraint(player *PlayerInGame, constraints []Constraint, costs []Cost) bool {
	fmt.Println("CONSTRAINT CHECK ", constraints, player.Nick)
	for _, c := range costs {
		switch op := c.Type; op {
		case "money":
			if player.Economy.Money < c.Value {
				fmt.Println("FAIL CONSTRAINT COST", c, player.Economy.Money)
				return false
			}
		case "science":
			if player.Civilian.NbResearchPoint < c.Value {
				fmt.Println("FAIL CONSTRAINT COST", c, player.Civilian.NbResearchPoint)
				return false
			}
		case "manpower":
			if player.Civilian.NbManpower < c.Value {
				fmt.Println("FAIL CONSTRAINT COST", c, player.Civilian.NbManpower)
				return false
			}
		case "morale":
			if player.Army.Morale < c.Value {
				fmt.Println("FAIL CONSTRAINT COST", c, player.Army.Morale)
				return false
			}

		}
	}
	for _, t := range constraints {
		if t.Type == "tech" && !StringInSlice(t.Value, player.PlayerTechnology) {
			fmt.Println("FAIL TECH CONSTRAINT PREREQUISITES", t, player.PlayerTechnology)
			return false
		}
	}

	return true

}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false

}

//ApplyEffect apply effect on player modifiers
func ApplyEffect(player *PlayerInGame, effect Effect) {
	for key, _ := range player.Modifiers {
		if key == effect.ModifierName {
			switch op := effect.Operator; op {
			case "+":
				player.Modifiers[key] += effect.Value
			case "-":
				player.Modifiers[key] -= effect.Value
			case "*":
				player.Modifiers[key] *= effect.Value
			case "/":
				player.Modifiers[key] *= 1 / effect.Value

			}
		}
	}
}

//ApplyCost apply cost on player board
func ApplyCost(player *PlayerInGame, cost Cost) {
	fmt.Println("APPLY COST !")
	switch op := cost.Type; op {
	case "money":
		player.Economy.Money -= cost.Value
	case "science":
		player.Civilian.NbResearchPoint -= cost.Value
	case "manpower":
		player.Civilian.NbManpower -= cost.Value
	case "morale":
		player.Army.Morale -= cost.Value

	}
}

func AlgoTerritorryChange(p1 *PlayerInGame, p2 *PlayerInGame, p1dmg float32, p2dmg float32) (*PlayerInGame, *PlayerInGame) {
	var winner, loser *PlayerInGame
	if p1dmg > (p2dmg * 1.05) {
		winner = p1
		loser = p2
	} else if p2dmg > (p1dmg * 1.05) {
		winner = p2
		loser = p1
	} else {
		return p1, p2
	}
	loser.Territory.Surface -= 0.01
	winner.Territory.Surface += 0.01

	return p1, p2
}

func AlgoRollTurnEvent(p1 *PlayerInGame, p2 *PlayerInGame, turn int) (*PlayerInGame, *PlayerInGame) {
	var allSingleEvents = GetEventsByType("Single")
	var currentEventp1 = GetEvent("event0")
	var currentEventp2 = GetEvent("event0")
	var totalWeight int
	for _, g := range allSingleEvents {
		totalWeight += g.Weight
	}
	totalWeight *= 10

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(totalWeight)
	for _, g := range allSingleEvents {
		r -= g.Weight
		if r <= 0 {
			currentEventp1 = g
			break
		}
	}
	for _, e := range currentEventp1.Effects {
		ApplyEffect(p1, e)
		p1.Logs = append(p1.Logs, PlayerLog{Turn: turn, ActionName: currentEventp1.ActionName})
	}
	r = rand.Intn(totalWeight)
	for _, g := range allSingleEvents {
		r -= g.Weight
		if r <= 0 {
			currentEventp2 = g
			break
		}
	}
	for _, e := range currentEventp2.Effects {
		ApplyEffect(p2, e)
		p2.Logs = append(p2.Logs, PlayerLog{Turn: turn, ActionName: currentEventp2.ActionName})
	}
	return p1, p2
}

//AlgoDamageDealt Calculate dmg dealt
func AlgoDamageDealt(player *PlayerInGame) float32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rollp1 = r1.Float32() + 4.0
	var rollp2 = r1.Float32() + 4.0
	var rollp3 = r1.Float32() + 4.0
	dmgModifier := (player.Army.Morale / 100.0) * (player.Army.Quality / 100.0)
	var dmgSoldier = player.Army.NbSoldier * 0.05 * rollp1 * player.Modifiers["soldierQuality"]
	var dmgLightTank = player.Army.NbLigtTank * 5 * rollp2 * player.Modifiers["lightTankQuality"]
	var dmgHvyTank = player.Army.NbHvyTank * 15 * rollp3 * player.Modifiers["heavyTankQuality"]
	var dmg = (dmgSoldier + dmgLightTank + dmgHvyTank) * 0.2 * dmgModifier
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
	var multiHvyTank float32
	var multiLgtTank float32
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
	armyUpkeep := (player.Army.NbSoldier * 100) + (player.Army.NbLigtTank * 1000) + (player.Army.NbHvyTank * 5000)
	tax := (player.Economy.TaxRate * 0.2 * player.Civilian.NbTotalCivil)
	player.Economy.Money = player.Economy.Money - armyUpkeep + tax

	//Technology

	player.Civilian.NbResearchPoint += player.Civilian.NbScientist * 0.05

	if player.Economy.Money > 0 {

		player.Army.NbLigtTank += player.Civilian.NbLightTankFactory * 3 * player.Modifiers["lightTankFactoryProduction"]
		player.Army.NbHvyTank += player.Civilian.NbHeavyTankFactory * 1 * player.Modifiers["heavyTankFactoryProduction"]

		player.Economy.Money -= (player.Civilian.NbLightTankFactory * 10000) + (player.Civilian.NbHeavyTankFactory * 100000)

		var nbThingToBuild float32 = 1.0

		var civilianProduction = player.Civilian.NbCivilianFactory * 0.01 * (2 / player.Economy.TaxRate) * (2 / player.ModifierPolicy.ManpowerSizePolicy)
		civilianProduction *= player.Modifiers["civilianFactoryProduction"]

		if player.ModifierPolicy.BuildLgtTankFac {
			nbThingToBuild += 1.0
		}
		if player.ModifierPolicy.BuildHvyTankFac {
			nbThingToBuild--
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
			player.Army.Morale--
		}
	}

	return player
}
