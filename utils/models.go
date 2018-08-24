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
	IsWar       bool
}

//ListPlayer list of players
type ListPlayer struct {
	player1 PlayerInGame
	player2 PlayerInGame
}

//PlayerInGame player ig
type PlayerInGame struct {
	PlayerID           int
	Nick               string
	Army               PlayerArmy
	ModifierPolicy     PlayerModifierPolicy
	Civilian           PlayerCivilian
	Economy            PlayerEconomy
	Territory          PlayerTerritory
	LastOrders         []PlayerLastOrders
	Technologies       []string
	Modifiers          map[string]float32
	Logs               []PlayerLog
	CallbackEffects    []CallbackEffect
	Policies           []PolicyValue
	PlayerInformations map[string]*PlayerInformation
	Country            Country
}

type Country struct {
	Name         string
	Effects      []Effect
	Description  string
	Flag         string
	Restrictions []Restriction
}

// type PlayerInformations struct {
// 	Type        string
// 	Description string
// 	Infos       []PlayerInformation
// }

type PlayerInformation struct {
	Type        string
	SubType     string
	Name        string
	Description string
	Value       float32
}

type PlayerLog struct {
	Turn       int
	ActionName string
}

//PlayerModifier modifeirs

type PlayerLastOrders struct {
	OrderID  string
	Cooldown int
}

//PlayerArmy current army of the player
type PlayerArmy struct {
	NbSoldier         float32
	NbLigtTank        float32
	NbHvyTank         float32
	NbArt             float32
	NbAirSup          float32
	NbAirBomb         float32
	Morale            float32
	Quality           float32
	InfantryEquipment float32
	Ammunition        float32
}

type PlayerCivilian struct {
	NbManpower        float32
	NbResearchPoint   float32
	NbScientist       float32
	NbCivilianFactory float32
}

type PlayerEconomy struct {
	Money                       float32
	Loans                       float32
	TaxRate                     float32
	LightTankProduction         float32
	HeavyTankProduction         float32
	ArtilleryProduction         float32
	InfantryEquipmentProduction float32
	AmmunitionProduction        float32
	FighterProduction           float32
	BomberProduction            float32
	FactoryProduction           float32
}

//PlayerModifierPolicy list of modifier policy
type PlayerModifierPolicy struct {
	TrainingPolicy     float32
	ManpowerSizePolicy float32
	ArtOnFactory       float32
	BuildLgtTankFac    float32
	BuildHvyTankFac    float32
	CivilianProduction float32
	TankProduction     float32
	AirCraftProduction float32
}

type PlayerTerritory struct {
	Surface      float32
	SmallCities  float32
	MediumCities float32
	BigCities    float32
	Barracks     float32
}

//GameMsg msg send to the routine
type GameMsg struct {
	GameID   uuid.UUID
	Text     string
	PlayerID int
	Action   string
	Type     string
	Cooldown int
	Value    float32
	Effects  []Effect
	Costs    []Cost
}

type Translation struct {
	ActionName  string
	Language    string
	ShortName   string
	LongName    string
	Description string
	ToolTip     string
}

type DisplayInfoCat struct {
	Category string
	Infos    []DisplayInfos
}

type DisplayInfos struct {
	Name         string
	Type         string
	LowAlert     float32
	VeryLowAlert float32
	GrowthName   string
	GrowthType   string
	Display      int
}
type ServerInfos struct {
	Region         string
	PlayersOnline  int
	PlayersWaiting int
	IsOnline       bool
	OnGoingGames   int
}

type News struct {
	Title     string
	Date      string
	Paragrahs []string
}

type Configuration struct {
	Connection_String string
}
