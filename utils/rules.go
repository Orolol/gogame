package utils

import (
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
	ID    int
	Value string
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
