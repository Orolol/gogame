package utils

import (
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

//Token Authentication token
type Token struct {
	gorm.Model
	Token  string
	Status string
}

//GameHistory list of past game
type GameHistory struct {
	gorm.Model
	GameEndState Game
	Player1ID    int
	Player2ID    int
}
