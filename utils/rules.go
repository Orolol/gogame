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

type PolicyChange struct {
	ID       int
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

type PlayerActionOrderApi struct {
	ID       int
	Value    float32
	PlayerID int
	GameID   uuid.UUID
}

type PlayerActionOrder struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	ActionName     string
	ConstraintName string
	Description    string
	Cooldown       int
}

type Technology struct {
	gorm.Model
	Name           string `gorm:"not null;unique"`
	Description    string
	TypeTechnology string
	Tier           int
	ActionName     string
	ConstraintName string
}
