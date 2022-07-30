package database

import (
	"api/types"

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

func CreateUser(u *types.User, t string) error {
	err := db.Create(u).Error
	if err != nil {
		return err
	}
	err = db.Create(&types.Token{Token: t, UserID: u.ID}).Error
	return err
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

func GetUserByID(id uint) *types.User {
	user := &types.User{}
	user.ID = id
	db.First(user)
	return user
}
