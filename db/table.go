package db

import (
	"fmt"
	"sync"
)

// Table represents a simple table with columns and rows
type Table struct {
	columns []Column
	rows    [][]interface{}
	mutex   sync.Mutex
}

// InsertRow inserts a new row into the table
func (t *Table) InsertRow(row []interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(row) != len(t.columns) {
		fmt.Println("Error: Row length does not match column length")
		return
	}
	t.rows = append(t.rows, row)
}

// SelectAll selects all rows from the table
func (t *Table) SelectAll() [][]interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.rows
}
