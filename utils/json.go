package utils

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)


func HandleReqJson(r *http.Request) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func HandleTableToJSON(columns []string, rows *sql.Rows) ([]map[string]interface{}, error) {
	var results []map[string]interface{} = []map[string]interface{}{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columns {
			columnPointers[i] = &values[i]
		}

		err := rows.Scan(columnPointers...)
		if err != nil {
			log.Fatal("Error in internal.GetAllAsk: ", err)
			return []map[string]interface{}{}, err
		}

		for i, column := range columns {
			val := values[i]
			row[column] = val
		}

		results = append(results, row)
	}

	return results, nil
}

