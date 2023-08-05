package internal

import (
	"fmt"
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

func (af *AsksFor) CreateAsksFor(what string, reason string, requester string, reviewer string) error {
	af.UUID = utils.UUID()
	af.CreatedTime = utils.CurrentTime()
	af.Status = constants.ASK_PENDING
	af.Reason = reason
	af.Reviewer = reviewer
	af.Requester = requester

	//FIXME validate data
	query := `INSERT INTO asks_for (uuid, what, created_time, modified_time, reason, status, requester, reviewer)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := storage.DatabaseExec(query, af.UUID, af.What, af.CreatedTime, af.CreatedTime, af.Reason, af.Status, af.Requester, af.Reviewer)
	if err != nil {
		utils.Logger("err", err.Error())
		return err
	}
	return nil
}

func (af *AsksFor) GetAsksFor(uuid string) *AsksFor {
	query := fmt.Sprintf(`SELECT * FROM asks_for WHERE uuid='%s'`, uuid)
	rows, err := storage.DatabaseQuery(query)
	if err != nil {
		utils.Logger("err", err.Error())
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		utils.Logger("err", err.Error())
		return nil
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		utils.Logger("err", err.Error())
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
		utils.Logger("err", err.Error())
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		utils.Logger("err", err.Error())
		return nil
	}

	results, err := utils.HandleTableToJSON(columns, rows)
	if err != nil {
		utils.Logger("err", err.Error())
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
