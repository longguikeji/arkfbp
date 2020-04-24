package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

// Table represents a table in the database
type Table struct {
	Name        string    `json:"name"`
	Type        TableType `json:"type"`
	Description string    `json:"description"`
}

// TableType ...
type TableType string

const (
	// AbstractTable marks this table will not be mapped a normal database in Disk
	AbstractTable TableType = "abstract"

	// StandardTable marks the table is a fair common table
	StandardTable TableType = "standard"
)

// GetTableTypeFromString ...
func GetTableTypeFromString(t string) (TableType, error) {
	switch t {
	case "abstract":
		return AbstractTable, nil

	case "standard":
		return StandardTable, nil
	}

	return "", errors.New("unknown table type")

}

// CreateTable creates a new table
func CreateTable(db *DB, name string, t string, description string) error {
	tableType, err := GetTableTypeFromString(t)
	if err != nil {
		return err
	}

	table := Table{
		Name:        name,
		Type:        tableType,
		Description: description,
	}

	data, err := json.Marshal(table)
	if err != nil {
		return err
	}

	tablePath := path.Join(db.Path(), "tables", name)

	_, err = os.Stat(tablePath)
	if err == nil {
		return errors.New("table exists")
	}

	os.MkdirAll(tablePath, os.ModePerm)

	tableDefinitionPath := path.Join(tablePath, "index.json")
	if err := ioutil.WriteFile(tableDefinitionPath, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// ListTables returns all the tables in the database
func ListTables(db *DB) ([]*Table, error) {
	tablePath := path.Join(db.Path(), "tables")

	dir, err := ioutil.ReadDir(tablePath)
	if err != nil {
		return nil, err
	}

	var tables []*Table

	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}

		var (
			definitionPath string
			table          Table
			data           []byte
			err            error
		)

		definitionPath = path.Join(tablePath, fi.Name(), "index.json")
		if _, err := os.Stat(definitionPath); err != nil {
			continue
		}

		data, err = ioutil.ReadFile(definitionPath)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &table); err != nil {
			return nil, err
		}

		tables = append(tables, &table)
	}

	return tables, nil
}
