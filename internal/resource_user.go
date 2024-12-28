package internal

import (
	"fmt"
	"time"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type ResourceUser struct {
	UUID         string
	UserID       string
	ResourceID   string
	Active       bool
	CreatedTime  string
	ModifiedTime string
}

func (ru *ResourceUser) CreateResourceUser(user_id string, resource_id string, active bool) (error, string) {
	existingResource := ru.GetResourceUser(user_id, resource_id)
	if existingResource != nil {
		return fmt.Errorf("resource_user already exists"), ""
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	ru.UUID = utils.UUID()
	ru.UserID = user_id
	ru.ResourceID = resource_id
	ru.Active = active
	ru.CreatedTime = createdTime
	ru.ModifiedTime = createdTime

	query := `
		INSERT INTO resource_user (uuid, user_id, resource_id, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, ru.UUID, ru.UserID, ru.ResourceID, ru.Active, ru.CreatedTime, ru.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createresourceuser: %v", err), ""
	}

	return nil, ru.UUID
}

func (ru *ResourceUser) GetResourceUser(user_id string, resource_id string) *ResourceUser {
	query := fmt.Sprintf(`SELECT * FROM resource_user WHERE user_id='%s' AND resource_id='%s'`, user_id, resource_id)

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

	return &ResourceUser{
		UUID:         results[0]["uuid"].(string),
		UserID:       results[0]["user_id"].(string),
		ResourceID:   results[0]["resource_id"].(string),
		Active:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (ru *ResourceUser) GetAllResourceUsers() []*ResourceUser {
	// TODO Add pagination
	query := `SELECT * FROM resource_user`

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

	resource_users := make([]*ResourceUser, 0, len(results))
	for _, res := range results {
		resource_user := &ResourceUser{
			UUID:         res["uuid"].(string),
			UserID:       res["user_id"].(string),
			ResourceID:   res["resource_id"].(string),
			Active:       res["active"].(bool),
			CreatedTime:  res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime: res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		resource_users = append(resource_users, resource_user)
	}

	return resource_users
}
