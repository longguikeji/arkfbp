package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// DB represents an ArkFBP Database
type DB struct {
	Name        string `json:"name"`
	Engine      string `json:"engine"`
	Description string `json:"description"`

	path string
}

// Path returns the db definition path
func (db *DB) Path() string {
	return db.path
}

// UnifyDatabaseName ...
func UnifyDatabaseName(name string) string {
	name = strings.ToLower(name)
	return name
}

// Create ...
func Create(root string, name string, engine string, description string) error {
	var (
		dbName           = UnifyDatabaseName(name)
		dbPath           = path.Join(root, dbName)
		dbDefinitionPath = path.Join(dbPath, "index.json")
		db               = DB{
			Name:        dbName,
			Engine:      engine,
			Description: description,
		}
	)

	// _, err := os.Stat(dbPath)
	// if err == nil {
	// 	fmt.Println(root, dbName, dbPath)
	// 	return errors.New("db already exists")
	// }

	// db
	if _, err := os.Stat(dbPath); err != nil {
		os.MkdirAll(dbPath, os.ModePerm)
	}

	// db/tables
	if _, err := os.Stat(path.Join(dbPath, "tables")); err != nil {
		os.MkdirAll(path.Join(dbPath, "tables"), os.ModePerm)
	}

	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dbDefinitionPath, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GenerateDSL ...
func GenerateDSL(dbPath string) error {
	return nil
}

// List returns all the databases
func List(root string) ([]*DB, error) {
	dir, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var databases []*DB

	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}

		var (
			dbDefinitionPath string
			db               DB
			data             []byte
			err              error
		)

		dbDefinitionPath = path.Join(root, fi.Name(), "index.json")
		if _, err := os.Stat(dbDefinitionPath); err != nil {
			continue
		}

		data, err = ioutil.ReadFile(dbDefinitionPath)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &db); err != nil {
			return nil, err
		}

		db.path = path.Join(root, fi.Name())

		databases = append(databases, &db)
	}

	return databases, nil
}

// GetByName return the DB found in terms of the name
func GetByName(root string, name string) (*DB, error) {
	dbPath := path.Join(root, name)
	if _, err := os.Stat(dbPath); err != nil {
		return nil, errors.New("db not found")
	}

	dbDefinitionPath := path.Join(dbPath, "index.json")
	if _, err := os.Stat(dbDefinitionPath); err != nil {
		return nil, errors.New("db not found")
	}

	var db DB
	data, err := ioutil.ReadFile(dbDefinitionPath)
	if err != nil {
		return nil, errors.New("failed to read the database information")
	}

	if err := json.Unmarshal(data, &db); err != nil {
		return nil, errors.New("failed to read the database information")
	}

	db.path = dbPath
	return &db, nil
}
