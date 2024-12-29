package internal

import (
	"fmt"
	"time"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type ResourceCredential struct {
	UUID         string
	CredentialID string
	ResourceID   string
	Active       bool
	CreatedTime  string
	ModifiedTime string
}

func (rc *ResourceCredential) CreateResourceCredential(credential_id string, resource_id string, active bool) (error, string) {
	existingResource := rc.GetResourceCredential(credential_id, resource_id)
	if existingResource != nil {
		return fmt.Errorf("resource_credential already exists"), ""
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	rc.UUID = utils.UUID()
	rc.CredentialID = credential_id
	rc.ResourceID = resource_id
	rc.Active = active
	rc.CreatedTime = createdTime
	rc.ModifiedTime = createdTime

	query := `
		INSERT INTO resource_credential (uuid, credential_id, resource_id, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, rc.UUID, rc.CredentialID, rc.ResourceID, rc.Active, rc.CreatedTime, rc.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createresourceuser: %v", err), ""
	}

	return nil, rc.UUID
}

func (rc *ResourceCredential) GetResourceCredential(credential_id string, resource_id string) *ResourceCredential {
	query := fmt.Sprintf(`SELECT * FROM resource_credential WHERE credential_id='%s' AND resource_id='%s'`, credential_id, resource_id)

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

	return &ResourceCredential{
		UUID:         results[0]["uuid"].(string),
		CredentialID: results[0]["credential_id"].(string),
		ResourceID:   results[0]["resource_id"].(string),
		Active:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (rc *ResourceCredential) GetAllResourceCredentials() []*ResourceCredential {
	// TODO Add pagination
	query := `SELECT * FROM resource_credential`

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

	resource_credentials := make([]*ResourceCredential, 0, len(results))
	for _, res := range results {
		resource_credential := &ResourceCredential{
			UUID:         res["uuid"].(string),
			CredentialID: res["credential_id"].(string),
			ResourceID:   res["resource_id"].(string),
			Active:       res["active"].(bool),
			CreatedTime:  res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime: res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		resource_credentials = append(resource_credentials, resource_credential)
	}

	return resource_credentials
}
