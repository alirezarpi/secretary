package internal

import (
	"bufio"
	"math/rand"
	"net"
	"fmt"
	"time"
	"strings"

	"secretary/alpha/utils"
	"secretary/alpha/storage"
)

const MIN = 1
const MAX = 100

type Session struct {
	UUID				 string
	UserID				 string
	ResourceCredentialID string
	TTL					 string
	ProxyURL			 string
	Active               bool
	CreatedTime			 string
	ModifiedTime		 string
}

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func HandleTCPConnection(c net.Conn) {
    utils.Logger("info", "serving TCP " + c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		// TODO check user permissions
		if err != nil {
		    utils.Logger("error", err.Error())
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		// TODO Audit here or in PostgreSQL Query
		result := "TARGET RESPONSE\n"
		c.Write([]byte(string(result)))
	    // TODO Create session record here.
	}
	c.Close()
}

func (s *Session) CreateSession(user_id string, resource_credential_id string, ttl string, active bool) (error, string) {

	// FIXME Add validation code here ...
	// FIXME change the error handling

	createdTime := utils.CurrentTime()

	s.UUID = utils.UUID()
	s.UserID = user_id
	s.ResourceCredentialID = resource_credential_id
	s.TTL = ttl
	s.Active = active
	s.CreatedTime = createdTime
	s.ModifiedTime = createdTime

	query := `
		INSERT INTO session (uuid, user_id, resouce_credential_id, ttl, proxy_url, active, created_time, modified_time)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := storage.DatabaseExec(query, s.UUID, s.UserID, s.ResourceCredentialID, s.TTL, s.ProxyURL, s.Active, s.CreatedTime, s.ModifiedTime)
	if err != nil {
		return fmt.Errorf("error in createsession: %v", err), ""
	}

	return nil, s.UUID
}

func (s *Session) GetSession(session_id string) *Session {
	query := fmt.Sprintf(`SELECT * FROM session WHERE uuid='%s'`, session_id)

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

	return &Session{
		UUID:				  results[0]["uuid"].(string),
		UserID:				  results[0]["user_id"].(string),
		ResourceCredentialID: results[0]["resource_credential_id"].(string),
		TTL:				  results[0]["ttl"].(string),
		ProxyURL:			  results[0]["proxy_url"].(string),
		Active:               results[0]["active"].(bool),
		CreatedTime:		  results[0]["created_time"].(time.Time).Format(time.RFC3339),
		ModifiedTime:		  results[0]["modified_time"].(time.Time).Format(time.RFC3339),
	}
}

