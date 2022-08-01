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
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	Username  string         `json:"username" gorm:"unique"`
	Bio       string         `json:"bio"`
	CreatedAt time.Time      `json:"joined"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}

type Message struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Text      string         `json:"text"`
	ToUser    uint           `json:"to_user"`
	FromUser  uint           `json:"from_user"`
	Viewed    time.Time      `json:"viewed"`
	CreatedAt time.Time      `json:"sended"`
	UpdatedAt time.Time      `json:"edited"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}
