package v1

import (
	"fmt"
	"math/rand"
	"net/http"
	"server/database"
	"server/errors"
	"server/types"
	"server/util"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := &types.NewUser{}
	err := c.BindJSON(user)
	if err != nil {
		util.BadRequest(c, errors.FailedToParseJSON)
		return
	} else if user.Name == "" {
		util.BadRequest(c, errors.UserNameEmpty)
		return
	} else if user.Username == "" {
		uname := ""
		for i := 0; i < 8; i++ {
			uname = uname + fmt.Sprintf("%c", rand.Intn(9))
		}
		user.Username = uname
	}
	token := util.GenToken()
	for database.TokenExists(token) {
		token = util.GenToken()
	}
	created_user, err := database.CreateUser(user, token)
	if err != nil {
		c.JSON(http.StatusConflict, errors.UserAlreadyExists)
		return
	}
	c.JSON(http.StatusOK, created_user)
}
