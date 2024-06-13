package database

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