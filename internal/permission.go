package internal

import (
	"fmt"
	"time"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type Permission struct {
	UUID			string
	Name			string
	Active			bool
	CreatedTime		string
	ModifiedTime	string
}

func (p *Permission) CreatePermission(name string, active bool) error {
	existingPerm := p.GetPermission(name)
	if existingPerm != nil {
		return fmt.Errorf("permission name %v already exists", name)
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	p.UUID = utils.UUID()
	p.Name = name
	p.CreatedTime = createdTime
	p.ModifiedTime = createdTime
	p.Active = active

	query := `
		INSERT INTO rbac_permission (uuid, name, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, p.UUID, p.Name, p.Active, p.CreatedTime, p.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createpermission: %v", err)
	}

	return nil
}

func (p *Permission) GetPermission(name string) *Permission {
	query := fmt.Sprintf(`SELECT * FROM rbac_permission WHERE name='%s'`, name)

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

	return &Permission{
		UUID:			results[0]["uuid"].(string),
		Name:			results[0]["name"].(string),
		Active:			results[0]["active"].(bool),
		CreatedTime:	results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime:	results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (p *Permission) GetAllPermissions() []*Permission {
	query := `SELECT * FROM rbac_permission`

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

	perms := make([]*Permission, 0, len(results))
	for _, res := range results {
		perm := &Permission{
			UUID:			res["uuid"].(string),
			Name:			res["name"].(string),
			Active:			res["active"].(bool),
			CreatedTime:	res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime:	res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		perms = append(perms, perm)
	}

	return perms
}
