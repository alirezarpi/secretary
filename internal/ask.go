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
		return false
	}
	return true
}

func (af *AsksFor) GetAsksFor(uuid string) *AsksFor {
	query := fmt.Sprintf(`SELECT * FROM asks_for WHERE uuid='%s'`, uuid[0])
	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetAsksFor: ", err)
		return nil
	}

	return results
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
