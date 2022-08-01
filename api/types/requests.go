package types

type FindUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type NewUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
}
