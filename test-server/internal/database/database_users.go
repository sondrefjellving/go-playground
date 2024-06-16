package database

import (
	"encoding/json"
	"errors"
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
		return User{}, err
	}

	id := getCorrectId(data.Users) 
	user := User{
		Id: id,
	}
	err = json.Unmarshal(payload, &user)
	if err != nil {
		return User{}, err
	}

	data.Users[id] = user
	if err := db.writeDB(data); err != nil {
		return User{}, nil
	}
	return user, nil
}

func (db *DB) DeleteUserById(id int) error {
	data, err := db.readDB()
	if err != nil {
		return err
	}

	if _, exists := data.Users[id]; !exists {
		return errors.New("cannot delete entry because it does not exist")
	}
	
	delete(data.Users, id)
	if err := db.writeDB(data); err != nil {
		return err
	}
	return nil
}

func getCorrectId(users map[int]User) int {
	id := len(users)
	for {
		if _, exists := users[id]; !exists {
			return id
		}
		id++
	}
}