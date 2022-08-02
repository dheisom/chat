package types

import "time"

type CreatedUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Token    string `json:"token"`
}

type SendedMessage struct {
	ID     uint   `json:"id"`
	Text   string `json:"text"`
	ToUser uint   `json:"to_user"`
}

type MessageResponse struct {
	ID       uint      `json:"id"`
	Text     string    `json:"text"`
	FromUser uint      `json:"from_user"`
	Sended   time.Time `json:"sended"`
	Viewed   bool      `json:"viewed"`
}
