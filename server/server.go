package server

import (
	"bufio"
	"encoding/json"
	"net"
	"strings"
)

var users = map[string]string{
	"admin": "password", // usuario: contraseña
}

// HandleConnection handles a single connection to the server
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}

		var request map[string]interface{}
		if err := json.Unmarshal([]byte(message), &request); err != nil {
			conn.Write([]byte("Error: Invalid request format\n"))
			continue
		}

		user, ok := request["user"].(string)
		password, ok2 := request["password"].(string)
		if !ok || !ok2 || users[user] != password {
			conn.Write([]byte("Error: Invalid username or password\n"))
			continue
		}

		handleRequest(conn, request)
	}
}
