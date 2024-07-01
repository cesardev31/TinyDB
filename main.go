package main

import (
	"fmt"
	"net"

	"github.com/cesardev31/TinyDB/db"
	"github.com/cesardev31/TinyDB/server"
)

func main() {
	// Load existing tables from disk
	db.LoadTable("test_table")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 8080...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go server.HandleConnection(conn)
	}
}
