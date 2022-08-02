package database

import (
	"server/types"

	"gorm.io/gorm"
)

var db *gorm.DB

func Start(d gorm.Dialector) {
	var err error
	db, err = gorm.Open(d)
	if err != nil {
		panic("Failed to start database: " + err.Error())
	}
	RunMigrations()
}

func RunMigrations() {
	db.AutoMigrate(&types.Token{})
	db.AutoMigrate(&types.User{})
	db.AutoMigrate(&types.Message{})
}

func CreateUser(u *types.NewUser, t string) (*types.CreatedUser, error) {
	user := &types.User{
		Name:     u.Name,
		Username: u.Username,
		Bio:      u.Bio,
	}
	err := db.Create(user).Error
	if err != nil {
		return &types.CreatedUser{}, err
	}
	err = db.Create(&types.Token{Token: t, UserID: user.ID}).Error
	created_user := &types.CreatedUser{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Bio:      user.Bio,
		Token:    t,
	}
	return created_user, err
}

func SaveMessage(m *types.Message) (uint, error) {
	m.ID = 0
	err := db.Create(m).Error
	return m.ID, err
}

func TokenExists(t string) bool {
	token := &types.Token{}
	db.Find(token, &types.Token{Token: t})
	return token.ID != 0
}

func GetUserByToken(t string) *types.User {
	user := &types.User{}
	db.Joins(
		"JOIN tokens ON tokens.user_id = users.id AND tokens.token = ?",
		t,
	).Find(user)
	return user
}

func GetUser(u *types.User) (*types.User, error) {
	user := &types.User{}
	err := db.First(user, u).Error
	return user, err
}

func GetMessages(u uint, s uint, l int) []types.Message {
	var messages []types.Message
	db.Where("messages.id >= ?", s).Limit(l).Find(
		&messages,
		&types.Message{FromUser: u},
	)
	return messages
}

func GetChats(u uint) []types.Chat {
	var users []types.User
	db.Joins(
		"JOIN messages ON (messages.from_user=users.id OR messages.to_user=users.id) AND users.id=?",
		u,
	).Distinct().Find(&users)
	var chats []types.Chat
	for i := range users {
		chats = append(chats, types.Chat{
			ID:    users[i].ID,
			Title: users[i].Name,
		})
	}
	return chats
}
