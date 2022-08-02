package v1

import (
	"net/http"
	"server/database"
	"server/errors"
	"server/types"
	"server/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMe(c *gin.Context) {
	token := c.Param("token")
	me := database.GetUserByToken(token)
	c.JSON(http.StatusOK, me)
}

func GetUser(c *gin.Context) {
	if c.Query("user") == "" {
		util.BadRequest(c, errors.What)
		return
	}
	identifier := &types.User{}
	user_param := c.Query("user")
	if id, err := strconv.Atoi(user_param); err == nil {
		if id <= 0 {
			util.BadRequest(c, errors.InvalidID)
			return
		}
		identifier.ID = uint(id)
	} else {
		identifier.Username = user_param
	}
	user, err := database.GetUser(identifier)
	if err == gorm.ErrRecordNotFound {
		util.BadRequest(c, errors.UserNotExists)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unknow error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func SendMessage(c *gin.Context) {
	token := c.Param("token")
	received_message := &types.NewMessage{}
	err := c.BindJSON(received_message)
	if err != nil {
		util.BadRequest(c, errors.NewE("failed to parse JSON: "+err.Error()))
		return
	}
	if received_message.Text == "" {
		util.BadRequest(c, errors.EmptyMessage)
		return
	} else if received_message.ToUser == "" {
		util.BadRequest(c, errors.NoDestinatary)
		return
	}
	user_identifier := &types.User{}
	if id, err := strconv.Atoi(received_message.ToUser); err == nil {
		if id <= 0 {
			util.BadRequest(c, errors.InvalidID)
			return
		}
		user_identifier.ID = uint(id)
	} else {
		user_identifier.Username = received_message.ToUser
	}
	to_user, err := database.GetUser(user_identifier)
	if err != nil {
		util.BadRequest(c, errors.ReceiverNotExists)
		return
	}
	sender := database.GetUserByToken(token)
	message := &types.Message{
		Text:     received_message.Text,
		FromUser: sender.ID,
		ToUser:   to_user.ID,
	}
	database.SaveMessage(message)
	c.JSON(http.StatusOK, types.SendedMessage{
		ID:     message.ID,
		Text:   message.Text,
		ToUser: message.ToUser,
	})
}

func GetMessages(c *gin.Context) {
	token := c.Param("token")
	uid := database.GetUserByToken(token).ID
	var seek uint = 1
	if c.Query("seek") != "" {
		s, err := strconv.Atoi(c.Query("seek"))
		if err != nil || s <= 0 {
			util.BadRequest(c, errors.InvalidSeek)
			return
		}
		seek = uint(s)
	}
	messages := database.GetMessages(uid, seek)
	var response []types.MessageResponse
	for i := range messages {
		if messages[i].Text == "" {
			continue
		}
		response = append(response, types.MessageResponse{
			ID:       messages[i].ID,
			Text:     messages[i].Text,
			FromUser: messages[i].FromUser,
			Sended:   messages[i].CreatedAt,
			Viewed:   messages[i].Viewed != messages[i].CreatedAt,
		})
	}
	c.JSON(http.StatusOK, response)
}
