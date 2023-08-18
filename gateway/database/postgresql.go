package database

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/lib/pq"
)

// -----------------------------------------
//			THIS IS PLACEHOLDER
// -----------------------------------------

const (
	proxyHost     = "0.0.0.0"
	proxyPort     = "4321"
	dbHost        = "your_database_host"
	dbPort        = 5432
	dbUser        = "your_database_user"
	dbPassword    = "your_database_password"
	dbName        = "your_database_name"
)

func main() {
	ln, err := net.Listen("tcp", proxyHost+":"+proxyPort)
	if err != nil {
		fmt.Println("Error starting proxy server:", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Proxy server listening on", proxyHost+":"+proxyPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(client net.Conn) {
	defer client.Close()

	fmt.Printf("Accepted connection from %s\n", client.RemoteAddr())

	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := pq.Open(dbConnStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	buf := make([]byte, 4096)
	for {
		n, err := client.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			break
		}

		command := string(buf[:n])
		fmt.Printf("Received command: %s\n", command)

		_, err = db.Exec(command)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

		response := []byte("Command received and logged.\n")
		_, err = client.Write(response)
		if err != nil {
			fmt.Println("Error writing response to client:", err)
			break
		}
	}
	fmt.Printf("Connection from %s closed\n", client.RemoteAddr())
}

