package utils

import (
	"github.com/google/uuid"
)

//Account Account Model
type Policy struct {
	Name          string `gorm:"not null;unique"`
	ActionName    string
	Constraints   []Constraint
	Description   string
	PossibleValue string
	TypePolicy    string
	DefaultValue  string
}

//API type for policy
type PolicyChange struct {
	ID       string
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

//API type for action
type PlayerActionOrderApi struct {
	ID       string
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

//SQL type for Actions
type PlayerActionOrder struct {
	Name        string `gorm:"not null;unique"`
	ActionName  string
	Constraints []Constraint
	Description string
	Costs       []Cost
	Cooldown    int
}

//Technology SQL type for technology
type Technology struct {
	Name           string `gorm:"not null;unique"`
	Description    string
	TypeTechnology string
	Tier           int
	Costs          []Cost
	ActionName     string
	Constraints    []Constraint
	Effects        []Effect
}

type Effect struct {
	ModifierName string
	Operator     string
	Value        float32
}

type Cost struct {
	Type  string
	Value float32
}

type Constraint struct {
	Type  string
	Value string
}

type PlayerEvent struct {
	Type        string
	Description string
	Constraints []Constraint
	Effects     []Effect
	ActionName  string
	Name        string
	Weight      int
}

// //Constraint Json type for constraint
// type Constraint struct {
// 	Tech   []string `json:tech`
// 	Turn   int      `json:turn`
// 	Policy []string `json:policy`
// 	Action []string `json:action`
// }
