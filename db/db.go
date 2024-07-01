package db

// CreateTable creates a new table with specified columns
func CreateTable(name string, columns []Column) *Table {
	table := &Table{
		columns: columns,
		rows:    [][]interface{}{},
	}
	Tables[name] = table
	SaveTable(name, table)
	return table
}
