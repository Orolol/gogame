package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//CheckConstraint Check if constraint is respected.
func CheckConstraint(player *PlayerInGame, constraints []Constraint, costs []Cost, game *Game) bool {
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
		if t.Type == "tech" && !StringInSlice(t.Value, player.Technologies) {
			return false
		} else if t.Type == "turn" {
			turn, _ := strconv.Atoi(t.Value)
			return CheckOperator(float32(turn), t.Operator, float32(game.CurrentTurn))
		} else if t.Type == "isWar" && !game.IsWar {
			return false
		} else if t.Type == "isNotWar" && game.IsWar {
			return false
		} else if t.Type == "Modifier" {
			value, _ := strconv.Atoi(t.Value)
			for key := range player.Modifiers {
				if key == t.Key {
					return CheckOperator(float32(value), t.Operator, float32(player.Modifiers[key]))
				}
			}
			return false
		} else if t.Type == "ModifierTurn" {
			for key := range player.Modifiers {
				if key == t.Key {
					return CheckOperator(float32(player.Modifiers[key]), t.Operator, float32(game.CurrentTurn))
				}
			}
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
func ApplyEffect(player *PlayerInGame, effect Effect, game *Game) {

	if effect.ModifierType == "Army" {
		switch field := effect.ModifierName; field {
		case "Morale":
			player.Army.Morale = ApplyOperator(effect.Value, effect.Operator, player.Army.Morale, game)
		case "Quality":
			player.Army.Quality = ApplyOperator(effect.Value, effect.Operator, player.Army.Quality, game)
		case "NbHvyTank":
			player.Army.NbHvyTank = ApplyOperator(effect.Value, effect.Operator, player.Army.NbHvyTank, game)
		case "NbLigtTank":
			player.Army.NbLigtTank = ApplyOperator(effect.Value, effect.Operator, player.Army.NbLigtTank, game)
		case "NbArt":
			player.Army.NbArt = ApplyOperator(effect.Value, effect.Operator, player.Army.NbArt, game)
		case "NbSoldier":
			player.Army.NbSoldier = ApplyOperator(effect.Value, effect.Operator, player.Army.NbSoldier, game)
		}

	} else if effect.ModifierType == "Economy" {
		switch field := effect.ModifierName; field {
		case "Money":
			player.Economy.Money = ApplyOperator(effect.Value, effect.Operator, player.Economy.Money, game)
		case "TaxRate":
			player.Economy.TaxRate = ApplyOperator(effect.Value, effect.Operator, player.Economy.Money, game)
		}
	} else if effect.ModifierType == "Civilian" {
		switch field := effect.ModifierName; field {
		case "NbManpower":
			player.Civilian.NbManpower = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbManpower, game)
		case "NbScientist":
			player.Civilian.NbScientist = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbScientist, game)
		case "NbTotalCivil":
			player.Civilian.NbTotalCivil = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbTotalCivil, game)
		case "NbResearchPoint":
			player.Civilian.NbResearchPoint = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbResearchPoint, game)
		case "NbLightTankFactory":
			player.Civilian.NbLightTankFactory = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbLightTankFactory, game)
		case "NbHeavyTankFactory":
			player.Civilian.NbHeavyTankFactory = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbHeavyTankFactory, game)
		case "NbCivilianFactory":
			player.Civilian.NbCivilianFactory = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbCivilianFactory, game)
		}
	} else if effect.ModifierType == "Policy" {
		switch field := effect.ModifierName; field {
		case "AirCraftProduction":
			player.ModifierPolicy.AirCraftProduction = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.AirCraftProduction, game)
		case "ArtOnFactory":
			player.ModifierPolicy.ArtOnFactory = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.ArtOnFactory, game)
		case "BuildHvyTankFac":
			player.ModifierPolicy.BuildHvyTankFac = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.BuildHvyTankFac, game)
		case "BuildLgtTankFac":
			player.ModifierPolicy.BuildLgtTankFac = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.BuildLgtTankFac, game)
		case "CivilianProduction":
			player.ModifierPolicy.CivilianProduction = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.CivilianProduction, game)
		case "ManpowerSizePolicy":
			player.ModifierPolicy.ManpowerSizePolicy = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.ManpowerSizePolicy, game)
		case "RecruitmentPolicy":
			player.ModifierPolicy.RecruitmentPolicy = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.RecruitmentPolicy, game)
		case "TankProduction":
			player.ModifierPolicy.TankProduction = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.TankProduction, game)
		}
	} else {
		player.Modifiers[effect.ModifierName] = ApplyOperator(effect.Value, effect.Operator, player.Modifiers[effect.ModifierName], game)

	}

	for _, cb := range effect.Callbacks {
		player.CallbackEffects = append(player.CallbackEffects, cb)
	}

}

func ApplyOperator(value float32, operator string, baseValue float32, game *Game) float32 {
	switch op := operator; op {
	case "+":
		return baseValue + value
	case "-":
		return baseValue - value
	case "*":
		return baseValue * value
	case "/":
		return baseValue / value
	case "turn+":
		return float32(game.CurrentTurn) + value
	case "=":
		return value
	default:
		return baseValue

	}
}
func CheckOperator(value float32, operator string, baseValue float32) bool {
	switch op := operator; op {
	case ">":
		return baseValue > value
	case "<":
		return baseValue < value
	case "=":
		return baseValue == value
	}
	return false
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

func AlgoRollTurnEvent(p1 *PlayerInGame, game *Game) *PlayerInGame {
	var allSingleEvents = GetEventsByType("Single")
	var currentEventp1 = GetEvent("event0")
	var totalWeight int
	var allPossibleEents []PlayerEvent
	for _, g := range allSingleEvents {
		if CheckConstraint(p1, g.Constraints, nil, game) {
			totalWeight += g.Weight
			allPossibleEents = append(allPossibleEents, g)
		}

	}
	totalWeight *= 10

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(totalWeight)
	for _, g := range allPossibleEents {
		r -= g.Weight
		if r <= 0 {
			currentEventp1 = g
			break
		}
	}
	for _, e := range currentEventp1.Effects {
		ApplyEffect(p1, e, game)
		p1.Logs = append(p1.Logs, PlayerLog{Turn: game.CurrentTurn, ActionName: currentEventp1.ActionName})
	}

	return p1
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
	natGrowth := player.ModifierPolicy.ManpowerSizePolicy * 0.001 * player.Civilian.NbTotalCivil
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

		if player.ModifierPolicy.BuildLgtTankFac == 1. {
			nbThingToBuild += 1.0
		}
		if player.ModifierPolicy.BuildHvyTankFac == 1. {
			nbThingToBuild--
		}
		if player.ModifierPolicy.BuildLgtTankFac == 1. {
			player.Civilian.NbLightTankFactory += (civilianProduction / nbThingToBuild) * 0.5
		}
		if player.ModifierPolicy.BuildHvyTankFac == 1. {
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
