package errors

type E map[string]string

var (
	// Users
	TokenNotExists    = E{"error": "this token does not exists"}
	UserNameEmpty     = E{"error": "the user name is empty"}
	UserAlreadyExists = E{"error": "this user already exists"}
	UserNotExists     = E{"error": "this user not exists"}
	InvalidID         = E{"error": "this ID is invalid"}

	// Messages
	EmptyMessage  = E{"error": "the text of the message is empty"}
	NoDestinatary = E{"error": "your message is without destinatary"}
	What          = E{"error": "what are you trying to do?"}

	// Data
	FailedToParseJSON = E{"error": "failed to parse JSON data"}
)
