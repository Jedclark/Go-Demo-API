package models

import "time"

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(id string, username string, password string, createdAt time.Time) User {
	return User{
		Id:        id,
		Username:  username,
		Password:  password,
		CreatedAt: createdAt,
	}
}

func (user User) ValidateUser() bool {
	if user.Username == "" || user.Password == "" {
		return false
	}
	return true
}
