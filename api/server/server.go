package server

import (
	"api/database"
	"api/errors"
	"api/types"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Start(p string) {
	server := gin.Default()
	api := server.Group("/api")
	{
		api.POST("/createUser", CreateUser)
		methods := api.Group("/:token")
		{
			methods.POST("/sendMessage", SendMessage)
			methods.GET("/getMe", GetMe)
		}
	}
	err := server.Run(":" + p)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}

func badRequest(c *gin.Context, e gin.H) {
	c.JSON(http.StatusBadRequest, e)
}

func genToken() string {
	var token strings.Builder
	lowerCharSet := "abcdedfghijklmnopqrst"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberSet := "0123456789"
	allCharSet := lowerCharSet + upperCharSet + numberSet
	for i := 0; i < 16; i++ {
		n := rand.Intn(len(allCharSet))
		token.WriteString(string(allCharSet[n]))
	}
	inRune := []rune(token.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func CreateUser(c *gin.Context) {
	user := &types.User{}
	err := c.BindJSON(user)
	if err != nil {
		badRequest(c, errors.FailedToParseJSON)
		return
	}
	if user.Name == "" {
		badRequest(c, errors.UserNameEmpty)
		return
	} else if user.Username == "" {
		uname := ""
		for i := 0; i < 8; i++ {
			uname = uname + string(rand.Intn(9))
		}
		user.Username = uname
	}
	token := genToken()
	for database.TokenExists(token) {
		token = genToken()
	}
	err = database.CreateUser(user, token)
	if err != nil {
		c.JSON(http.StatusConflict, errors.UserAlreadyExists)
		return
	}
	c.JSON(http.StatusOK, types.CreatedUser{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	})
}

func GetMe(c *gin.Context) {
	token := c.Param("token")
	if !database.TokenExists(token) {
		badRequest(c, errors.TokenNotExists)
		return
	}
	me := database.GetUserByToken(token)
	c.JSON(http.StatusOK, me)
}

func SendMessage(c *gin.Context) {
	token := c.Param("token")
	if !database.TokenExists(token) {
		badRequest(c, errors.TokenNotExists)
		return
	}
	message := &types.Message{}
	err := c.BindJSON(message)
	if err != nil {
		badRequest(c, gin.H{
			"error": "failed to parse JSON: " + err.Error(),
		})
		return
	}
	if message.Text == "" {
		badRequest(c, errors.EmptyMessage)
		return
	} else if message.ToUser == 0 {
		badRequest(c, errors.NoDestinatary)
		return
	}
	sender := database.GetUserByToken(token)
	message.FromUser = sender.ID
	database.SaveMessage(message)
	c.JSON(http.StatusOK, types.SendedMessage{
		ID:       message.ID,
		Text:     message.Text,
		FromUser: sender.ID,
		ToUser:   message.ToUser,
	})
}
