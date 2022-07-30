package types

type CreatedUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SendedMessage struct {
	ID       uint   `json:"id"`
	Text     string `json:"text"`
	FromUser uint   `json:"from"`
	ToUser   uint   `json:"to"`
}
