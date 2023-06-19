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
		log.Fatal("Error in GetAllAsk: ", err)
		return []map[string]interface{}{}
	}

	var results []map[string]interface{} = []map[string]interface{}{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columns {
			columnPointers[i] = &values[i]
		}

		err = rows.Scan(columnPointers...)
		if err != nil {
			log.Fatal("Error in internal.GetAllAsk: ", err)
			return []map[string]interface{}{}
		}

		for i, column := range columns {
			val := values[i]
			row[column] = val
		}

		results = append(results, row)
	}

	return results
}

func GetAsk(uuid string) map[string]interface{} {
	query := fmt.Sprintf(`SELECT * from asks_for where uuid=%s`, uuid)
	rows, _ := storage.DatabaseQuery(query)
	return map[string]interface{}{
		"result": rows,
	}
}

