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
	GameID   uuid.UUID
	Text     string
	PlayerID int
	Action   string
	Value    float32
}
