package main

import (
	"encoding/json"
	"os"
	"strings"
)

type UserRepo struct {
}

const (
	filename = "user.txt"
)

type User struct {
	ChatID int64  `json:"chatID"`
	State  string `json:"state"`
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (g *UserRepo) GetUser() (*User, error) {
	user := User{}
	file, err := os.ReadFile(filename)
	if err != nil {
		if strings.Contains(err.Error(), "system cannot find the file specified") {
			return nil, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &user)

	return &user, err
}

func (g *UserRepo) SetUser(user User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, os.ModePerm)
}
