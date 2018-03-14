package utils

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Account Account Model
type Policy struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	ActionName     string
	ConstraintName string
	Description    string
	PossibleValue  string
	TypePolicy     string
	DefaultValue   string
}

//API type for policy
type PolicyChange struct {
	ID       int
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

//API type for action
type PlayerActionOrderApi struct {
	ID       int
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

//SQL type for Actions
type PlayerActionOrder struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	ActionName     string
	ConstraintName string
	Description    string
	Cost           float32
	Cooldown       int
}

//Technology SQL type for technology
type Technology struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Description    string
	TypeTechnology string
	Tier           int
	Cost           float32
	ActionName     string
	ConstraintName string
}

//Constraint Json type for constraint
type Constraint struct {
	Tech   []string `json:tech`
	Turn   int      `json:turn`
	policy []string `json:policy`
	action []string `json:action`
}
