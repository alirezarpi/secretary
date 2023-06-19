package internal

import (
	"fmt"
	"log"

	"secretary/alpha/internal/constants"
	"secretary/alpha/storage"
	"secretary/alpha/utils"
)


func CreateAsk(what string, reason string) bool {
	uuid := utils.UUID()
	createdTime := utils.CurrentTime()
	status := constants.ASK_PENDING

	query := fmt.Sprintf(`INSERT INTO asks_for (uuid, what, created_time, modified_time, reason, status)
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s')`, uuid, what, createdTime, createdTime, reason, status)

	log.Printf("Asking query: %s", query)

	_, err := storage.DatabaseExec(query)
	if err != nil {
		return false
	}
	return true
}

func GetAllAsk() []map[string]interface{} {
	query := fmt.Sprintf(`SELECT * from asks_for`)

	rows, err := storage.DatabaseQuery(query)
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("error in getallask: ", err)
		return []map[string]interface{}{}
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("error in getallask: ", err)
		return []map[string]interface{}{}
	}

	return results
}

func GetAsk(uuid string) []map[string]interface{} {
	query := fmt.Sprintf(`SELECT * from asks_for where uuid='%s'`, uuid)
	
	rows, err := storage.DatabaseQuery(query)
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("error in getallask: ", err)
		return []map[string]interface{}{}
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("error in getallask: ", err)
		return []map[string]interface{}{}
	}

	return results
}

