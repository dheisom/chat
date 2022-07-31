package errors

type E map[string]string

func NewE(d string) E {
	return E{"error": d}
}

var (
	// Users
	TokenNotExists    = NewE("this token does not exists")
	UserNameEmpty     = NewE("the user name is empty")
	UserAlreadyExists = NewE("this user already exists")
	UserNotExists     = NewE("this user not exists")
	InvalidID         = NewE("this ID is invalid")

	// Messages
	EmptyMessage  = NewE("the text of the message is empty")
	NoDestinatary = NewE("your message is without destinatary")
	What          = NewE("what are you trying to do?")

	// Data
	FailedToParseJSON = NewE("failed to parse JSON data")
)
