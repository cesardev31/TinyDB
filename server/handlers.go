package server

import (
	"encoding/json"
	"net"

	"github.com/cesardev31/TinyDB/db"
)

// handleRequest maneja las solicitudes según la acción especificada.
func handleRequest(conn net.Conn, request map[string]interface{}) {
	switch request["action"] {
	case "create_table":
		name := request["name"].(string)
		columns := request["columns"].([]interface{})
		columnDefs := make([]db.Column, len(columns))
		for i, col := range columns {
			colMap := col.(map[string]interface{})
			columnDefs[i] = db.Column{
				Name: colMap["name"].(string),
				Type: db.DataType(colMap["type"].(float64)),
			}
		}
		db.CreateTable(name, columnDefs)
		conn.Write([]byte("Table created\n"))

	case "insert_row":
		name := request["name"].(string)
		row := request["row"].([]interface{})
		if table, exists := db.Tables[name]; exists {
			table.InsertRow(row)
			db.SaveTable(name, table)
			conn.Write([]byte("Row inserted\n"))
		} else {
			conn.Write([]byte("Error: Table not found\n"))
		}

	case "select_all":
		name := request["name"].(string)
		if table, exists := db.Tables[name]; exists {
			rows := table.SelectAll()
			response, _ := json.Marshal(rows)
			conn.Write([]byte(string(response) + "\n"))
		} else {
			conn.Write([]byte("Error: Table not found\n"))
		}

	default:
		conn.Write([]byte("Error: Unknown action\n"))
	}
}
