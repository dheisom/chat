package v1

import (
	"server/database"
	"server/errors"
	"server/util"

	"github.com/gin-gonic/gin"
)

func Registry(g *gin.RouterGroup) {
	g.POST("/createUser", CreateUser)
	methods := g.Group("/:token")
	methods.Use(CheckTokenMiddleware())
	{
		methods.POST("/sendMessage", SendMessage)
		methods.GET("/getMe", GetMe)
		methods.GET("/getUser", GetUser)
		methods.GET("/getMessages", GetMessages)
		methods.GET("/getChats", GetChats)
	}
}

func CheckTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		if !database.TokenExists(token) {
			util.BadRequest(c, errors.TokenNotExists)
			c.Abort()
		} else {
			c.Next()
		}
	}
}
