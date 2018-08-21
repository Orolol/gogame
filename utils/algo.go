package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//CheckRestriction Check if restrictions is respected.
func CheckRestrictionInGame(player *PlayerInGame, restrictions []Restriction, game *Game) bool {
	for _, r := range restrictions {

		switch t := r.Type; t {
		case "country":

			if player.Country.Name != r.Value {
				return false
			}
		}
	}
	return true
}

//CheckRestriction Check if restrictions is respected.
func CheckRestrictionBefore(acc *Account, restrictions []Restriction) bool {
	for _, r := range restrictions {
		fmt.Println('r', r)
		switch t := r.Type; t {
		case "country":
			fmt.Println("COUNTRY CHECK", acc.SelectedCountry, r.Value)
			if acc.SelectedCountry != r.Value {
				return false
			}
		}
	}
	return true
}

//CheckConstraint Check if constraint is respected.
func CheckConstraint(player *PlayerInGame, constraints []Constraint, costs []Cost, game *Game, valueToCheck float32) bool {
	//fmt.Println("CONSTRAINT CHECK ", constraints, player.Nick)
	for _, c := range costs {
		switch op := c.Type; op {
		case "money":
			if player.Economy.Money < c.Value {
				//fmt.Println("FAIL CONSTRAINT COST", c, player.Economy.Money)
				return false
			}
		case "science":
			if player.Civilian.NbResearchPoint < c.Value {
				//fmt.Println("FAIL CONSTRAINT COST", c, player.Civilian.NbResearchPoint)
				return false
			}
		case "manpower":
			if player.Civilian.NbManpower < c.Value {
				//fmt.Println("FAIL CONSTRAINT COST", c, player.Civilian.NbManpower)
				return false
			}
		case "morale":
			if player.Army.Morale < c.Value {
				//fmt.Println("FAIL CONSTRAINT COST", c, player.Army.Morale)
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
		} else if t.Type == "Custom" {
			value, _ := strconv.Atoi(t.Value)
			return CheckOperator(float32(value), t.Operator, valueToCheck)
		} else if t.Type == "linked" {
			if t.Value == "production" {
				if player.Economy.HeavyTankProduction+player.Economy.AmmunitionProduction+player.Economy.LightTankProduction+player.Economy.FactoryProduction+player.Economy.ArtilleryProduction+player.Economy.InfantryEquipmentProduction+player.Economy.BomberProduction+player.Economy.FighterProduction > 800 {
					fmt.Println("LINKED CONSTRAINT FAIL", player.Economy)
					return false
				}
			}
			return true
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
	fmt.Println(player.Nick, "APPLY EFFECT", effect)
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
		case "Loans":
			player.Economy.Loans = ApplyOperator(effect.Value, effect.Operator, player.Economy.Loans, game)
		case "TaxRate":
			player.Economy.TaxRate = ApplyOperator(effect.Value, effect.Operator, player.Economy.Money, game)
		case "ArtilleryProduction":
			player.Economy.ArtilleryProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.ArtilleryProduction, game)
		case "BomberProduction":
			player.Economy.BomberProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.BomberProduction, game)
		case "FighterProduction":
			player.Economy.FighterProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.FighterProduction, game)
		case "HeavyTankProduction":
			player.Economy.HeavyTankProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.HeavyTankProduction, game)
		case "InfantryEquipmentProduction":
			player.Economy.InfantryEquipmentProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.InfantryEquipmentProduction, game)
		case "LightTankProduction":
			player.Economy.LightTankProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.LightTankProduction, game)
		case "AmmunitionProduction":
			player.Economy.AmmunitionProduction = ApplyOperator(effect.Value, effect.Operator, player.Economy.AmmunitionProduction, game)
		}
	} else if effect.ModifierType == "Civilian" {
		switch field := effect.ModifierName; field {
		case "NbManpower":
			player.Civilian.NbManpower = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbManpower, game)
		case "NbScientist":
			player.Civilian.NbScientist = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbScientist, game)
		case "NbResearchPoint":
			player.Civilian.NbResearchPoint = ApplyOperator(effect.Value, effect.Operator, player.Civilian.NbResearchPoint, game)
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
		case "TrainingPolicy":
			player.ModifierPolicy.TrainingPolicy = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.TrainingPolicy, game)
		case "TankProduction":
			player.ModifierPolicy.TankProduction = ApplyOperator(effect.Value, effect.Operator, player.ModifierPolicy.TankProduction, game)
		}
	} else if effect.ModifierType == "Territory" {
		switch field := effect.ModifierName; field {
		case "Barracks":
			player.Territory.Barracks = ApplyOperator(effect.Value, effect.Operator, player.Territory.Barracks, game)
		}
	} else {
		player.Modifiers[effect.ModifierName] = ApplyOperator(effect.Value, effect.Operator, player.Modifiers[effect.ModifierName], game)

	}

	for _, cb := range effect.Callbacks {
		player.CallbackEffects = append(player.CallbackEffects, cb)
	}

	fmt.Println(player.Nick, "APPLIED EFFECT", effect)
	//fmt.Println("#############################################")

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
	////fmt.Println("APPLY COST !")
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
	loser.Territory.Surface--
	winner.Territory.Surface++

	loser.Territory.SmallCities--
	winner.Territory.SmallCities++

	if (loser.Territory.Surface < 90 || winner.Territory.Surface < 90) && int(loser.Territory.Surface)%2 == 0 {
		loser.Territory.Barracks--
		winner.Territory.Barracks++
	}

	if (loser.Territory.Surface < 85 || winner.Territory.Surface < 85) && int(loser.Territory.Surface)%4 == 0 {
		loser.Territory.MediumCities--
		winner.Territory.MediumCities++
	}
	if (loser.Territory.Surface < 50 || winner.Territory.Surface < 50) && int(loser.Territory.Surface)%10 == 0 {
		loser.Territory.BigCities--
		winner.Territory.BigCities++
	}

	return p1, p2
}

func AlgoRollTurnEvent(p1 *PlayerInGame, game *Game) *PlayerInGame {
	var allSingleEvents = GetEventsByType("Single")
	var currentEventp1 = GetEvent("event0")
	var totalWeight int
	var allPossibleEents []PlayerEvent
	for _, g := range allSingleEvents {
		if CheckConstraint(p1, g.Constraints, nil, game, 0) {
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
	}
	if currentEventp1.ActionName != "" {
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
	var dmgSoldier = player.Army.NbSoldier * 0.05 * rollp1 * (player.Modifiers["soldierQuality"] / 100)
	var dmgLightTank = player.Army.NbLigtTank * 5 * rollp2 * (player.Modifiers["lightTankQuality"] / 100)
	var dmgHvyTank = player.Army.NbHvyTank * 15 * rollp3 * (player.Modifiers["heavyTankQuality"] / 100)
	var dmgAerial = AlgoAerialBomb(player) * player.Modifiers["bomberTargetArmy"]
	var dmg = (dmgSoldier + dmgLightTank + dmgHvyTank + dmgAerial) * 0.2 * dmgModifier

	player.PlayerInformations["SoldierDmg"].Value = dmgSoldier
	player.PlayerInformations["LightTankDmg"].Value = dmgLightTank
	player.PlayerInformations["HeavyTankDmg"].Value = dmgHvyTank
	player.PlayerInformations["TotalGroundDmg"].Value = dmg
	////fmt.Println("DAMAGE ", dmg)
	return dmg
}

func AlgoDamageDealtOnFactories(player *PlayerInGame) float32 {
	dmgModifier := (player.Army.Morale / 100.0) * (player.Army.Quality / 100.0)
	var dmgAerial = AlgoAerialBomb(player) * player.Modifiers["bomberTargetFactory"]
	var dmg = (dmgAerial) * 0.05 * dmgModifier
	return dmg
}

func AlgoDamageDealtOnPopulation(player *PlayerInGame) float32 {
	dmgModifier := (player.Army.Morale / 100.0) * (player.Army.Quality / 100.0)
	var dmgAerial = AlgoAerialBomb(player) * player.Modifiers["bomberTargetPopulation"]
	var dmg = (dmgAerial) * 0.05 * dmgModifier
	return dmg
}

func AlgoAerialCombat(player *PlayerInGame) float32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rollp1 = r1.Float32() + 4.0

	dmgModifier := (player.Army.Morale / 100.0) * (player.Army.Quality / 100.0)

	var dmgFighter = player.Army.NbAirSup * 10 * rollp1 * player.Modifiers["engageFighter"] * (player.Modifiers["dmgAerialBonus"] / 100)
	var dmgBomber = player.Army.NbAirBomb * 2 * rollp1 * player.Modifiers["engageBomber"] * (player.Modifiers["dmgAerialBonus"] / 100)

	player.PlayerInformations["AirSupAerialDmg"].Value = dmgFighter
	player.PlayerInformations["AirBombAerialDmg"].Value = dmgBomber
	player.PlayerInformations["TotalAerialDmg"].Value = (dmgFighter + dmgBomber) * dmgModifier

	return (dmgFighter + dmgBomber) * dmgModifier
}

func AlgoAerialBomb(player *PlayerInGame) float32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rollp1 = r1.Float32() + 4.0

	var dmgFighter = player.Army.NbAirSup * 5 * rollp1 * player.Modifiers["engageFighter"]
	var dmgBomber = player.Army.NbAirBomb * 20 * rollp1 * player.Modifiers["engageBomber"]

	player.PlayerInformations["AirSupGroundDmg"].Value = dmgFighter * player.Modifiers["bomberTargetArmy"] * (player.Modifiers["dmgBombBonus"] / 100)
	player.PlayerInformations["AirBombGroundDmg"].Value = dmgBomber * player.Modifiers["bomberTargetArmy"] * (player.Modifiers["dmgBombBonus"] / 100)

	return dmgFighter + dmgBomber
}

func AlgoFullAerialPhase(p1 *PlayerInGame, p2 *PlayerInGame) (*PlayerInGame, *PlayerInGame) {
	var p1AD = AlgoAerialCombat(p1)
	var p2AD = AlgoAerialCombat(p2)

	var p1engBomb, p1engFight, p2engBomb, p2engFight float32

	p1engFight = p1.Army.NbAirSup * p1.Modifiers["engageFighter"]
	p1engBomb = p1.Army.NbAirBomb * p1.Modifiers["engageBomber"]
	p2engFight = p2.Army.NbAirSup * p2.Modifiers["engageFighter"]
	p2engBomb = p2.Army.NbAirBomb * p2.Modifiers["engageBomber"]

	// fmt.Println("AERIAL SUP P1", p1.Army.NbAirSup, p1AD)
	// fmt.Println("AERIAL SUP P2", p2.Army.NbAirSup, p2AD)

	// fmt.Println("AERIAL BOMB P1", p1.Army.NbAirBomb, p1AD)
	// fmt.Println("AERIAL BOMB P2", p2.Army.NbAirBomb, p2AD)

	var lossp1sup, lossp2sup, lossp2bomb, lossp1bomb float32

	if p2engBomb+p2engFight != 0 {
		lossp2sup = (p1AD / 120) * (p2engFight / (p2engBomb + p2engFight))
		lossp2bomb = (p1AD / 50) * (p2engBomb / (p2engBomb + p2engFight))
	}
	if p1engBomb+p1engFight != 0 {
		lossp1sup = (p2AD / 120) * (p1engFight / (p1engBomb + p1engFight))
		lossp1bomb = (p2AD / 50) * (p1engBomb / (p1engBomb + p1engFight))
	}

	if p2.Army.NbAirSup < lossp2sup {
		p2.Army.NbAirSup = 0
	} else {
		p2.Army.NbAirSup -= lossp2sup
	}

	if p1.Army.NbAirSup < lossp1sup {
		p1.Army.NbAirSup = 0
	} else {
		p1.Army.NbAirSup -= lossp1sup
	}

	if p2.Army.NbAirBomb < lossp2bomb {
		p2.Army.NbAirBomb = 0
	} else {
		p2.Army.NbAirBomb -= lossp2bomb
	}

	if p1.Army.NbAirBomb < lossp1bomb {
		p1.Army.NbAirBomb = 0
	} else {
		p1.Army.NbAirBomb -= lossp1bomb
	}

	// fmt.Println("AFTER AERIAL SUP P1", p1.Army.NbAirSup, p1AD, lossp1sup)
	// fmt.Println("AFTER AERIAL SUP P2", p2.Army.NbAirSup, p2AD, lossp2sup)

	// fmt.Println("AFTER AERIAL BOMB P1", p1.Army.NbAirBomb, p1AD, lossp1bomb)
	// fmt.Println("AFTER AERIAL BOMB P2", p2.Army.NbAirBomb, p2AD, lossp2bomb)

	return p1, p2

}

//AlgoReinforcement calc reinforcement
func AlgoReinforcement(player *PlayerInGame) *PlayerInGame {
	if player.Economy.Money > 0 {
		// var minRf float32 = 100.0
		// reinforcement := 1000 * player.ModifierPolicy.TrainingPolicy
		// ////fmt.Println("REINFORCEMENT : ", reinforcement)
		// if reinforcement > player.Civilian.NbManpower {
		// 	reinforcement = 0.0
		// } else if reinforcement < minRf {
		// 	reinforcement = minRf
		// }

		reinforcement := player.Territory.Barracks * 50 * player.ModifierPolicy.TrainingPolicy

		player.PlayerInformations["SoldierRecruit"].Value = reinforcement

		if player.Army.InfantryEquipment > reinforcement {
			player.Army.NbSoldier += reinforcement
			player.Civilian.NbManpower -= reinforcement
			player.Army.InfantryEquipment -= reinforcement
		} else {
			player.Army.NbSoldier += player.Army.InfantryEquipment
			player.Civilian.NbManpower -= player.Army.InfantryEquipment
			player.Army.InfantryEquipment = 0.
		}

	}
	natGrowth := player.Territory.Barracks * 50 * player.ModifierPolicy.ManpowerSizePolicy
	player.Civilian.NbManpower += natGrowth

	player.PlayerInformations["ManpowerGrowth"].Value = natGrowth

	return player
}

func AlgoGetPopulation(player *PlayerInGame) float32 {
	return (player.Territory.SmallCities * 20000) + (player.Territory.MediumCities * 100000) + (player.Territory.BigCities * 1000000) + 2000000
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
		//fmt.Println("Civilian damage ", dmgIncoming-totalHp)
	}
	//fmt.Println("DMG MODIFER", (player.Army.Morale/100.0)*(player.Army.Quality/100.0))
	dmgModifier := 2 / (1 + (player.Army.Morale/100.0)*(player.Army.Quality/100.0))
	player.Army.NbSoldier -= dmgIncoming * multiSoldier * 0.1 * dmgModifier
	player.Army.NbLigtTank -= dmgIncoming * multiLgtTank * 0.02 * dmgModifier
	player.Army.NbHvyTank -= dmgIncoming * multiHvyTank * 0.005 * dmgModifier

	if player.Army.NbSoldier < 0 {
		player.Army.NbSoldier = 0.0
	}
	if player.Army.NbLigtTank < 0 {
		//fmt.Println("All lght tank lost ")
		player.Army.NbLigtTank = 0.0
	}
	if player.Army.NbHvyTank < 0 {
		//fmt.Println("All hvy tank lost ")
		player.Army.NbHvyTank = 0.0
	}

	return player
}

//AlgoDamageRepartition Calculate loses
func AlgoDamageRepartitionOnFactories(player *PlayerInGame, dmgIncoming float32) *PlayerInGame {
	loss := (dmgIncoming / (3000))
	if player.Civilian.NbCivilianFactory-loss > 0 {
		player.Civilian.NbCivilianFactory -= loss
	} else {
		player.Civilian.NbCivilianFactory = 0
	}
	return player
}

func AlgoEconomicEndTurn(player *PlayerInGame) *PlayerInGame {

	// if player.Economy.Money < 0 {
	// 	if player.Army.Morale > 10 {
	// 		player.Army.Morale--
	// 	}
	// }

	if player.Economy.Loans > 0 {
		player.Economy.Money -= player.Economy.Loans * 1000000
	}
	armyUpkeep := (player.Army.NbSoldier * 25) + (player.Army.NbLigtTank * 500) + (player.Army.NbHvyTank * 2000)
	tax := (player.Economy.TaxRate * 0.35 * AlgoGetPopulation(player))
	player.Economy.Money = player.Economy.Money - armyUpkeep + tax
	player.PlayerInformations["TaxRevenu"].Value = tax
	player.PlayerInformations["TaxRate"].Value = player.Economy.TaxRate
	//Technology

	player.Civilian.NbResearchPoint += player.Civilian.NbScientist * 0.05 * player.Modifiers["researchEfficiency"]
	if player.Economy.Money > 0 {
		var civilianProduction = player.Civilian.NbCivilianFactory * 0.05 * (2 / player.Economy.TaxRate) * player.Modifiers["workersConcrptionEfficiency"]
		civilianProduction *= player.Modifiers["civilianFactoryProduction"]

		var litghTankProd = (player.Economy.LightTankProduction / 100) * 3 * player.Modifiers["lightTankFactoryProduction"] * civilianProduction
		var heavyTankProd = (player.Economy.HeavyTankProduction / 100) * 1 * player.Modifiers["heavyTankFactoryProduction"] * civilianProduction
		var bbProd = (player.Economy.BomberProduction / 100) * 0.5 * civilianProduction
		var fProd = (player.Economy.FighterProduction / 100) * 1 * civilianProduction
		var artProd = (player.Economy.ArtilleryProduction / 100) * 0.5 * civilianProduction
		var civProd = (player.Economy.FactoryProduction / 100) * 0.01 * civilianProduction
		var amuProd = (player.Economy.AmmunitionProduction / 100) * 100 * civilianProduction
		var infProd = (player.Economy.InfantryEquipmentProduction / 100) * 150 * civilianProduction

		player.Army.NbLigtTank += litghTankProd
		player.Army.NbHvyTank += heavyTankProd
		player.Army.NbAirBomb += bbProd
		player.Army.NbAirSup += fProd
		player.Army.NbArt += artProd
		player.Army.Ammunition += amuProd
		player.Army.InfantryEquipment += infProd
		player.Civilian.NbCivilianFactory += civProd

		player.PlayerInformations["lightTankProduction"].Value = litghTankProd
		player.PlayerInformations["heavyTankProduction"].Value = heavyTankProd
		player.PlayerInformations["bomberProduction"].Value = bbProd
		player.PlayerInformations["fighterProduction"].Value = fProd

		player.PlayerInformations["artilleryProduction"].Value = artProd
		player.PlayerInformations["factoryProduction"].Value = civProd
		player.PlayerInformations["infantryEquipmentProduction"].Value = infProd
		player.PlayerInformations["ammunitionProduction"].Value = amuProd

		player.PlayerInformations["Factory"].Value = player.Civilian.NbCivilianFactory

		// }
		// if player.Economy.Money > 0 {
		// 	var nbThingToBuild float32 = 1.0

		// 	player.PlayerInformations["civilianProduction"].Value = civilianProduction

		// 	if player.ModifierPolicy.BuildLgtTankFac == 1. {
		// 		nbThingToBuild += 1.0
		// 	}
		// 	if player.ModifierPolicy.BuildHvyTankFac == 1. {
		// 		nbThingToBuild--
		// 	}
		// 	if player.ModifierPolicy.BuildLgtTankFac == 1. {
		// 		player.Civilian.NbLightTankFactory += (civilianProduction / nbThingToBuild) * 0.1
		// 	}
		// 	if player.ModifierPolicy.BuildHvyTankFac == 1. {
		// 		player.Civilian.NbHeavyTankFactory += (civilianProduction / nbThingToBuild) * 0.05
		// 	}
		// 	player.Civilian.NbCivilianFactory += (civilianProduction / nbThingToBuild) * 0.15

		if player.Army.Morale > 100 {
			player.Army.Morale--
		} else {
			player.Army.Morale++
		}

	}
	return player
}
