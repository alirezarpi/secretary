package internal

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type DatabaseResource struct {
	UUID string
	Name string

	DBType         string
	DBNames        []string
	DBPort         int
	DBHost         string
	DBUser         string
	DBPasswordHash string

	Active       bool
	CreatedTime  string
	ModifiedTime string
}

func (r *DatabaseResource) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r.DBPasswordHash = string(hash)
	return nil
}

func (r *DatabaseResource) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(r.DBPasswordHash), []byte(password))
	return err == nil
}

func (r *DatabaseResource) CreateDatabaseResource(
	name string,
	active bool,
	dbType string,
	dbNames []string,
	dbPort int,
	dbHost string,
	dbUser string,
	dbPassword string) error {

	existingResource := r.GetDatabaseResource(name)
	if existingResource != nil {
		return fmt.Errorf("resource already exists")
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	r.UUID = utils.UUID()
	r.Name = name
	r.DBType = dbType
	r.DBNames = dbNames
	r.DBPort = dbPort
	r.DBHost = dbHost
	r.DBUser = dbUser
	r.SetPassword(dbPassword)
	r.Active = active
	r.CreatedTime = createdTime
	r.ModifiedTime = createdTime

	print("fdsfsdfsdfs")
	query := `
		INSERT INTO resource_ (uuid, name, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, r.UUID, r.Name, r.Active, r.CreatedTime, r.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createresource: %v", err)
	}

	return nil
}

func (r *DatabaseResource) GetDatabaseResource(name string) *DatabaseResource {
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

	return &DatabaseResource{
		UUID:         results[0]["uuid"].(string),
		Name:         results[0]["name"].(string),
		Active:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (r *DatabaseResource) GetAllDatabaseResources() []*DatabaseResource {
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

	resources := make([]*DatabaseResource, 0, len(results))
	for _, res := range results {
		resource := &DatabaseResource{
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
