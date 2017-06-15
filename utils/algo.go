package utils

import (
	"math/rand"
	"time"
)

//AlgoDamageDealt Calculate dmg dealt
func AlgoDamageDealt(player *PlayerInGame) float32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var rollp1 = r1.Float32()
	var dmg = player.Army.NbSoldier * 0.5 * rollp1
	return dmg
}

//AlgoReinforcement calc reinforcement
func AlgoReinforcement(player *PlayerInGame) float32 {
	var minRf float32 = 100.0
	reinforcement := player.NbPop * 0.01 * player.ModifierPolicy.RecruitmentPolicy
	if reinforcement < minRf {
		return minRf
	}
	return reinforcement
}

//AlgoDamageRepartition Calculate loses
func AlgoDamageRepartition(player *PlayerInGame, dmgIncoming float32) *PlayerInGame {
	player.Army.NbSoldier -= dmgIncoming
	return player
}
