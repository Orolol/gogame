package utils

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Account Account Model
type Account struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Login    string `gorm:"not null;unique"`
	Password string
	Token    Token
	TokenID  uint
	ELO      int
}

//GameHistory GameHistory Model
type GameHistory struct {
	gorm.Model
	Winner  uint
	Loser   uint
	GameID  uuid.UUID
	ELODiff int
}

//Token Authentication token
type Token struct {
	gorm.Model
	Token  string
	Status string
}

//GameHistory list of past game
