package internal

import (
	"fmt"
	"time"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type ResourceDatabase struct {
	UUID           string
	Name		   string
	ResourceUserID string
	DatabaseHost   string
	DatabasePort   string
	Active         bool
	CreatedTime    string
	ModifiedTime   string
}

func (rd *ResourceDatabase) CreateResourceDatabase(name string, resource_user_id string, db_host string, db_port string, active bool) (error, string) {
	existingResource := rd.GetResourceDatabase(resource_user_id, db_host, db_port)
	if existingResource != nil {
		return fmt.Errorf("resource_database already exists"), ""
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	rd.UUID = utils.UUID()
	rd.Name = name
	rd.ResourceUserID = resource_user_id
	rd.DatabaseHost = db_host
	rd.DatabasePort = db_port
	rd.Active = active
	rd.CreatedTime = createdTime
	rd.ModifiedTime = createdTime

	query := `
		INSERT INTO resource_database (uuid, name, resource_user_id, db_host, db_port, active, created_time, modified_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, rd.UUID, rd.Name, rd.ResourceUserID, rd.DatabaseHost, rd.DatabasePort, rd.Active, rd.CreatedTime, rd.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createdatabaseresource: %v", err), ""
	}

	return nil, rd.UUID
}

func (rd *ResourceDatabase) GetResourceDatabase(resource_user_id string, db_host string, db_port string) *ResourceDatabase {
	query := fmt.Sprintf(`SELECT * FROM resource_database WHERE resource_user_id='%s' AND db_host='%s' AND db_port='%s'`, resource_user_id, db_host, db_port)

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

	return &ResourceDatabase{
		UUID:			results[0]["uuid"].(string),
		ResourceUserID: results[0]["user_resource_id"].(string),
		DatabaseHost:   results[0]["db_host"].(string),
		DatabasePort:   results[0]["db_port"].(string),
		Active:			results[0]["active"].(bool),
		CreatedTime: 	results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime:	results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (rd *ResourceDatabase) GetAllResourceDatabases() []*ResourceDatabase {
	// TODO Add pagination
	query := `SELECT * FROM resource_database`

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

	resource_databases := make([]*ResourceDatabase, 0, len(results))
	for _, res := range results {
		resource_database := &ResourceDatabase{
			UUID:			res["uuid"].(string),
			ResourceUserID: res["user_resource_id"].(string),
			DatabaseHost:   res["db_host"].(string),
			DatabasePort:   res["db_port"].(string),
			Active:			res["active"].(bool),
			CreatedTime:	res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime:	res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		resource_databases = append(resource_databases, resource_database)
	}

	return resource_databases
}
