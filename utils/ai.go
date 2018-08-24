package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var AIactions []PlayerActionOrder
var AItech []Technology

func InitAIEasy() {
	actions := []string{"actionWarPropaganda", "promoteWarHearoes", "emergencyRecruitment"}
	techs := []string{"technoIndusT1N2", "technoIndusT2N2", "technoIndusT3N2", "technoMilT1N2", "technoMilT2N2", "technoMilT3N2"}
	for _, a := range GetActions() {
		if Contains(actions, a.ActionName) {
			AIactions = append(AIactions, a)
		}
	}
	for _, t := range GetTechnolgies() {
		if Contains(techs, t.ActionName) {
			AItech = append(AItech, t)
		}
	}
}

func RollAiAction(player *PlayerInGame, game *Game) {
	var totalWeight int

	var allPossible []interface{}
	var currentAction interface{}
	actionRoll := false
	for _, g := range AIactions {
		if CheckConstraint(player, g.Constraints, g.Costs, game, 0) {
			totalWeight += 1
			allPossible = append(allPossible, g)
		}

	}
	for _, t := range AItech {
		if CheckConstraint(player, t.Constraints, t.Costs, game, 0) && !Contains(player.Technologies, t.ActionName) {
			totalWeight += 1
			allPossible = append(allPossible, t)
		}

	}
	totalWeight *= 5

	if totalWeight != 0 {

		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(totalWeight)
		for _, g := range allPossible {
			r -= 1
			if r <= 0 {
				fmt.Println("ROLL !", g)
				currentAction = g
				actionRoll = true
				break
			}
		}
		if actionRoll {
			switch currentAction.(type) {
			case Technology:
				player.Technologies = append(player.Technologies, currentAction.(Technology).ActionName)
				fmt.Println("AI DID SOMETHING ACTION", currentAction, player.Nick)
				for _, e := range currentAction.(Technology).Effects {
					ApplyEffect(player, e, game)
				}
				for _, c := range currentAction.(Technology).Costs {
					ApplyCost(player, c)
				}

			case PlayerActionOrder:
				fmt.Println("AI DID SOMETHING TECH", currentAction, player.Nick)
				for _, e := range currentAction.(PlayerActionOrder).Effects {
					ApplyEffect(player, e, game)
				}
				for _, c := range currentAction.(PlayerActionOrder).Costs {
					ApplyCost(player, c)
				}

			}

		}

	}

}
