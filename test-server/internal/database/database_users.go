package database

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (db *DB) GetUsers() ([]User, error) {
	data, err := db.readDB()
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(data.Users))
	for _, user := range data.Users {
		users = append(users, user)
	}

	return users, nil
}

func (db *DB) GetUserById(id int) (User, error) {
	data, err := db.readDB()
	if err != nil {
		return User{}, err
	}
	user, exists := data.Users[id]
	if !exists {
		return User{}, errors.New("Cannot find user with id: " + string(id))
	}

	return user, nil
}

func (db *DB) CreateUser(payload []byte) (User, error) {
	data, err := db.readDB()
	if err != nil {
		fmt.Println("1")
		return User{}, err
	}

	id := len(data.Users)
	user := User{
		Id: id,
	}
	err = json.Unmarshal(payload, &user)
	if err != nil {
		fmt.Println("2")
		return User{}, err
	}

	data.Users[id] = user
	if err := db.writeDB(data); err != nil {
		fmt.Println("3")
		return User{}, nil
	}
	return user, nil
}