package errors

import (
	"github.com/gin-gonic/gin"
)

var (
	TokenNotExists    = gin.H{"error": "this token does not exists"}
	FailedToParseJSON = gin.H{"error": "failed to parse JSON data"}
	UserNameEmpty     = gin.H{"error": "the user name is empty"}
	UserAlreadyExists = gin.H{"error": "this user already exists"}
	EmptyMessage      = gin.H{"error": "the text of the message is empty"}
	NoDestinatary     = gin.H{"error": "your message is without destinatary"}
)
