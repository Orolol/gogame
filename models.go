package main

import "github.com/google/uuid"

//GameManager list of current game.
type GameManager struct {
	GameList map[uuid.UUID]Game
}

//GameConf Game conf
type GameConf struct {
	GameType  string
	NbPlayers int
}

//Game Game state
type Game struct {
	GameID      uuid.UUID
	CurrentTurn int
	ListPlayers []PlayerInGame
	Conf        GameConf
	Queue       chan GameMsg
}

//ListPlayer list of players
type ListPlayer struct {
	player1 PlayerInGame
	player2 PlayerInGame
}

//PlayerInGame player ig
type PlayerInGame struct {
	PlayerID       int
	nick           string
	Army           PlayerArmy
	NbPop          float32
	ModifierPolicy PlayerModifierPolicy
}

//PlayerArmy current army of the player
type PlayerArmy struct {
	NbSoldier  float32
	NbLigtTank float32
	NbHvyTank  float32
	NbArt      float32
	NbAirSup   float32
	NbAirBomb  float32
}

//PlayerModifierPolicy list of modifier policy
type PlayerModifierPolicy struct {
	RecruitmentPolicy float32
}

//GameMsg msg send to the routine
type GameMsg struct {
	GameID   int
	Text     string
	PlayerID int
	Action   PlayerAction
	value    float32
}
