package internal

import (
	"fmt"
	"log"

	"secretary/alpha/internal/constants"
	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type AsksFor struct {
	UUID         string
	What         string
	Reason       string
	Status       string
	Requester    string
	Reviewer     string
	CreatedTime  string
	ModifiedTime string
}

func (af *AsksFor) CreateAsksFor(what string, reason string) error {
	uuid := utils.UUID()
	createdTime := utils.CurrentTime()
	status := constants.ASK_PENDING
	requester := "placeholder"
	reviewer := "placeholder"

	query := `INSERT INTO asks_for (uuid, what, created_time, modified_time, reason, status, requester, reviewer)
		VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`

	_, err := storage.DatabaseExec(query, uuid, what, createdTime, createdTime, reason, status, requester, reviewer)
	if err != nil {
		log.Fatal("Error in CreateAsksFor: ", err)
		return err
	}

	return nil
}

func (af *AsksFor) GetAsksFor(uuid string) *AsksFor {
	query := fmt.Sprintf(`SELECT * FROM asks_for WHERE uuid='%s'`, uuid[0])
	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetAsksFor: ", err)
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error in GetAsksFor: ", err)
		return nil
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("Error in GetAsksFor: ", err)
		return nil
	}

	if len(results) == 0 {
		return nil
	}

	return &AsksFor{
		UUID:         results[0]["uuid"].(string),
		What:     results[0]["username"].(string),
		Reason: results[0]["password_hash"].(string),
		Status:       results[0]["active"].(bool),
		Requester:       results[0]["active"].(bool),
		Reviewer:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (af *AsksFor) GetAllAsksFors() *AsksFor {
	query = `SELECT * FROM asks_for`
	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetAllAskFor: ", err)
		return nil
	}

	return results
}
