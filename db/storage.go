package db

import (
	"encoding/json"
	"os"
)

var Tables = make(map[string]*Table)

// LoadTable loads a table from a file
func LoadTable(name string) (*Table, error) {
	file, err := os.Open(name + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			return &Table{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var table Table
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&table); err != nil {
		return nil, err
	}

	Tables[name] = &table
	return &table, nil
}

// SaveTable saves a table to a file
func SaveTable(name string, table *Table) error {
	file, err := os.Create(name + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(table); err != nil {
		return err
	}

	return nil
}
