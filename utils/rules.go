package utils

import (
	"github.com/google/uuid"
)

//Account Account Model
type Policy struct {
	Name           string `gorm:"not null;unique"`
	ActionName     string
	Constraints    []Constraint
	Description    string
	PossibleValue  string
	PossibleValue2 []PolicyValue
	TypePolicy     string
	DefaultValue   string
	MaxChange      float32
}

type PolicyValue struct {
	ActionName  string
	Value       float32
	Name        string
	Description string
	Constraints []Constraint
	Effects     []Effect
	IsDefault   bool
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
	Effects     []Effect
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
	ModifierType string
	ModifierName string
	Operator     string
	Value        float32
	Target       string
	Callbacks    []CallbackEffect
}

type CallbackEffect struct {
	Constraints []Constraint
	Effects     []Effect
}

type Cost struct {
	Type  string
	Value float32
}

type Constraint struct {
	Type     string
	Key      string
	Value    string
	Operator string
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
