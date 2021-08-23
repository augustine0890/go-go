package models

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
