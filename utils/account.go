package utils

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//Account Account Model
type Account struct {
	gorm.Model
	Name            string `gorm:"not null;unique"`
	Login           string `gorm:"not null;unique"`
	Password        string
	ProfilePic      string `gorm:"default:'pp1'"`
	Step            int    `gorm:"default:'0'"`
	TokenID         uint
	ELO             int
	SelectedCountry string
}

//Account Account Model
type AccountApi struct {
	ID              uint
	Login           string
	Name            string
	Token           string
	ELO             int
	ProfilePic      string
	Step            int
	SelectedCountry string
}

//Account Account Model
type GameHistoryApi struct {
	Created_at string
	WinnerNick string
	LoserNick  string
	ELODiff    int
}

//Account Account Model
type AccountLeaderboardApi struct {
	Name string
	ELO  int
}

//GameHistory GameHistory Model
type GameHistory struct {
	gorm.Model
	WinnerID uint
	LoserID  uint
	GameID   uuid.UUID
	ELODiff  int
}

//Token Authentication token
type Token struct {
	gorm.Model
	AccountID uint
	Token     string
	Status    string
}

type ProfilePic struct {
	Availablity string
	Name        string
	Directory   string
}

//GameHistory list of past game
