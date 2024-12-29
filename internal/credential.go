package internal

import (
	"fmt"
	"time"
	
	"golang.org/x/crypto/bcrypt"

	"secretary/alpha/storage"
	"secretary/alpha/utils"
)

type Credential struct {
	UUID         string
	Username	 string
	PasswordHash string
	Active       bool
	CreatedTime  string
	ModifiedTime string
}

func (c *Credential) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.PasswordHash = string(hash)
	return nil
}

func (c *Credential) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.PasswordHash), []byte(password))
	return err == nil
}

func (c *Credential) CreateCredential(username string, password string, active bool) (error, string) {
	//FIXME if same password with same username was passed it will allow it as it's checking with hash
	existingResource := c.GetCredential(username, password)
	if existingResource != nil {
		return fmt.Errorf("credential already exists"), ""
	}

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	c.UUID = utils.UUID()
	c.Username = username
	c.Active = active
	c.CreatedTime = createdTime
	c.ModifiedTime = createdTime

	err := c.SetPassword(password)
	if err != nil {
		return fmt.Errorf("err in setpassword: %v", err), ""
	}

	query := `
		INSERT INTO credential (uuid, username, password, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = storage.DatabaseExec(query, c.UUID, c.Username, c.PasswordHash, c.Active, c.CreatedTime, c.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createresourceuser: %v", err), ""
	}

	return nil, c.UUID
}

func (c *Credential) GetCredential(username string, password string) *Credential {
	query := fmt.Sprintf(`SELECT * FROM credential WHERE username='%s' AND password='%s'`, username, password)

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

	return &Credential{
		UUID:         results[0]["uuid"].(string),
		Username:	  results[0]["username"].(string),
		PasswordHash: results[0]["password"].(string),
		Active:       results[0]["active"].(bool),
		CreatedTime:  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime: results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

func (c *Credential) GetAllCredentials() []*Credential {
	// TODO Add pagination
	query := `SELECT * FROM credential`

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

	credentials := make([]*Credential, 0, len(results))
	for _, res := range results {
		credential := &Credential{
			UUID:         res["uuid"].(string),
			Username:	  res["username"].(string),
			PasswordHash: res["password"].(string),
			Active:       res["active"].(bool),
			CreatedTime:  res["created_time"].(time.Time).Format(time.RFC3339),
			ModifiedTime: res["modified_time"].(time.Time).Format(time.RFC3339),
		}
		credentials = append(credentials, credential)
	}

	return credentials
}
