package utils

import "github.com/google/uuid"

//GameManager list of current game.
type GameManager struct {
	GameList map[uuid.UUID]Game
}

//GameConf Game conf
type GameConf struct {
	GameType  string
	NbPlayers int
	Players   []Account
}

//Game Game state
type Game struct {
	GameID      uuid.UUID
	State       string
	Winner      PlayerInGame
	Loser       PlayerInGame
	CurrentTurn int
	ListPlayers []PlayerInGame
	Conf        GameConf
}

//ListPlayer list of players
type ListPlayer struct {
	player1 PlayerInGame
	player2 PlayerInGame
}

//PlayerInGame player ig
type PlayerInGame struct {
	PlayerID       int
	Nick           string
	Army           PlayerArmy
	ModifierPolicy PlayerModifierPolicy
	Civilian       PlayerCivilian
	Economy        PlayerEconomy
}

//PlayerArmy current army of the player
type PlayerArmy struct {
	NbSoldier  float32
	NbLigtTank float32
	NbHvyTank  float32
	NbArt      float32
	NbAirSup   float32
	NbAirBomb  float32
	Morale     float32
	Quality    float32
}

type PlayerCivilian struct {
	NbTotalCivil       float32
	NbManpower         float32
	NbCivilianFactory  float32
	NbLightTankFactory float32
	NbHeavyTankFactory float32
}

type PlayerEconomy struct {
	Money   float32
	TaxRate float32
}

//PlayerModifierPolicy list of modifier policy
type PlayerModifierPolicy struct {
	RecruitmentPolicy  float32
	ManpowerSizePolicy float32
	ArtOnFactory       bool
	BuildLgtTankFac    bool
	BuildHvyTankFac    bool
}

//GameMsg msg send to the routine
type GameMsg struct {
	GameID   uuid.UUID
	Text     string
	PlayerID int
	Action   string
	Value    float32
}
