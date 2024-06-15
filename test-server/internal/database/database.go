package database

import (
	"encoding/json"
	"os"
	"sync"
)

type DB struct {
	path string
	mux *sync.RWMutex
}

type DBStructure struct {
	Users map[int]User
}

type User struct {
	Id		int 	`json:"id"`
	Name 	string `json:"name"`
	Age 	int	   `json:"age"`	
}

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
		mux: &sync.RWMutex{},
	}

	err := db.setupDB()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) setupDB() error {
	if _, err := os.Open(db.path); err != os.ErrNotExist {
		return nil
	}
	file, err := os.Create(db.path)
	if err != nil {
		return err
	}
	defer file.Close()

	dbstruct := DBStructure{
		Users: make(map[int]User),
	}

	err = db.writeDB(dbstruct)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) writeDB(dbStruct DBStructure) error {
	dbAsJson, err := json.Marshal(dbStruct)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, dbAsJson, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) readDB() (DBStructure, error) {
	dbAsJson, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, db.setupDB()
	}

	dbStruct := DBStructure{}
	err = json.Unmarshal(dbAsJson, &dbStruct)
	if err != nil {
		return DBStructure{}, err
	}

	return dbStruct, nil
}