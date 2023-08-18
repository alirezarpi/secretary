package internal

import (
	"fmt"
	"time"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type Resource struct {
	UUID         string
	Name         string
	Active       bool
	CreatedTime  string
	ModifiedTime string
}

func (r *Resource) CreateResource(name string, active bool) error {
	existingResource := r.GetResource(name)
	if existingResource != nil {
		return fmt.Errorf("resource already exists")
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	r.UUID = utils.UUID()
	r.Name = name
	r.Active = active
	r.CreatedTime = createdTime
	r.ModifiedTime = createdTime

	query := `
		INSERT INTO resource (uuid, name, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, r.UUID, r.Name, r.Active, r.CreatedTime, r.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createresource: %v", err)
	}

	return nil
}

func (r *Resource) GetResource(name string) *Resource {
	query := fmt.Sprintf(`SELECT * FROM resource WHERE name='%s'`, name)

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

	return &Resource{
		UUID:         results[0]["uuid"].(string),
		Name:         results[0]["name"].(string),
		Active:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (r *Resource) GetAllResources() []*Resource {
	query := `SELECT * FROM resource`

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

	resources := make([]*Resource, 0, len(results))
	for _, res := range results {
		resource := &Resource{
			UUID:         res["uuid"].(string),
			Name:         res["name"].(string),
			Active:       res["active"].(bool),
			CreatedTime:  res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime: res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		resources = append(resources, resource)
	}

	return resources
}
