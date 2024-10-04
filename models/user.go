package models

import (
	"find-pairs/config"
)

type User struct {
	Username string
	Password string
}

func (u *User) Authenticate(username, password string) bool {
	return username == config.Username && password == config.Password
}
