package server

import (
	"api/database"
	"api/errors"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func badRequest(c *gin.Context, e errors.E) {
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

func isTokenOK(c *gin.Context) bool {
	token := c.Param("token")
	if !database.TokenExists(token) {
		badRequest(c, errors.TokenNotExists)
		return false
	}
	return true
}
