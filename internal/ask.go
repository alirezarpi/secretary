package internal

import (
	"fmt"
	"log"
	"time"

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
	query := fmt.Sprintf(`SELECT * FROM asks_for WHERE uuid='%s'`, uuid)
	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetAsksFor: ", err)
		return nil
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
		UUID:			results[0]["uuid"].(string),
		What:			results[0]["what"].(string),
		Reason:			results[0]["reason"].(string),
		Status:			results[0]["status"].(string),
		Requester:      results[0]["requester"].(string),
		Reviewer:       results[0]["reviewer"].(string),
		CreatedTime:	results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime:	results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (af *AsksFor) GetAllAsksFors() []*AsksFor {
	query := `SELECT * FROM asks_for`
	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		log.Fatal("Error in GetAllAskFor: ", err)
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("Error in GetAllAsksFors: ", err)
		return nil
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		log.Fatal("Error in GetAllAsksFors: ", err)
		return nil
	}


	asksFors := make([]*AsksFor, 0, len(results))
	for _, res := range results {
		asksFor := &AsksFor{
			UUID:			res["uuid"].(string),
			What:			res["what"].(string),
			Reason:			res["reason"].(string),
			Status:			res["status"].(string),
			Requester:      res["requester"].(string),
			Reviewer:       res["reviewer"].(string),
			CreatedTime:	res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime:	res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		asksFors = append(asksFors, asksFor)
	}
	return asksFors
}
