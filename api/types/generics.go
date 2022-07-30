package types

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	Token  string `gorm:"unique"`
	UserID uint   `gorm:"unique"`
	gorm.Model
}

type User struct {
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Bio      string `json:"bio"`
	gorm.Model
}

type Message struct {
	Text     string    `json:"text"`
	ToUser   uint      `json:"to"`
	FromUser uint      `json:"from"`
	Viewed   time.Time `json:"viewed"`
	gorm.Model
}
