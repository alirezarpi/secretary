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
	requester := "placeholder"
	reviewer := "placeholder"

	query := fmt.Sprintf(`INSERT INTO asks_for (uuid, what, created_time, modified_time, reason, status)
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`, uuid, what, createdTime, createdTime, reason, status, requester, reviewer)

	log.Printf("Asking query: %s", query)

	_, err := storage.DatabaseExec(query)
	if err != nil {
		return false
	}
	return true
}

func GetAsk(uuid ...string) []map[string]interface{} {
	var query string
	if len(uuid) > 0 {
		query = fmt.Sprintf(`SELECT * FROM asks_for WHERE uuid='%s'`, uuid[0])
	} else {
		query = `SELECT * FROM asks_for`
	}

	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetAsk: ", err)
		return []map[string]interface{}{}
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error in GetAsk: ", err)
		return []map[string]interface{}{}
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("Error in GetAsk: ", err)
		return []map[string]interface{}{}
	}

	return results
}
