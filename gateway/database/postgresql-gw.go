package database

import (
	"fmt"
	"io"
	"net"
	"secretary/alpha/utils"

	"github.com/lib/pq"
)

const (
	proxyHost  = "0.0.0.0"
	proxyPort  = "4321"
	dbHost     = "your_database_host"
	dbPort     = 5432
	dbUser     = "your_database_user"
	dbPassword = "your_database_password"
	dbName     = "your_database_name"
)

func PostgresqlDatabaseGateway(proxyPort string, proxyHost string) {
	ln, err := net.Listen("tcp", proxyHost+":"+proxyPort)
	if err != nil {
		utils.Logger("err", "[PG-DB-GW] could not start db listener: "+err.Error())
		return
	}
	defer ln.Close()
	utils.Logger("info", "[PG-DB-GW] Proxy server listening on "+proxyHost+":"+proxyPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			utils.Logger("err", "[PG-DB-GW] Error accepting connection: "+err.Error())
			continue
		}
		go handlePostgresqlDatabaseConnection(conn)
	}
}

func handlePostgresqlDatabaseConnection(client net.Conn) {
	defer client.Close()

	utils.Logger("debug", "[PG-DB-GW] Accepted connection from: "+client.RemoteAddr().String())

	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := pq.Open(dbConnStr)
	if err != nil {
		utils.Logger("err", "[PG-DB-GW] Error connecting to database:"+err.Error())
		return
	}
	defer db.Close()

	buf := make([]byte, 4096)
	for {
		n, err := client.Read(buf)
		if err != nil {
			if err != io.EOF {
				utils.Logger("error", "[PG-DB-GW] Error reading from client: "+err.Error())
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
