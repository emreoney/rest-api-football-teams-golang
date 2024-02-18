package models

import "gorm.io/gorm"

type Information struct {
	Message string `json:"message"`
}

type League struct {
	gorm.Model
	Name  string `json:"name"`
	Teams []Team `json:"teams" gorm:"foreignKey:LeagueID"`
}

type Team struct {
	gorm.Model
	Name     string   `json:"name"`
	LeagueID uint     `json:"leagueID"`
	Players  []Player `json:"players" gorm:"foreignKey:TeamID"`
}

type Player struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
	Position string `json:"position"`
	TeamID   uint   `json:"teamID"`
}
