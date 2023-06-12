package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"database/sql"

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

func HandleTableToJSON(dbPath, tableName string) ([]byte, error) {
	// FIXME create function that returns db obj
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM " + tableName

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	results := make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		rowData := make(map[string]interface{})

		for i, col := range values {
			rowData[columns[i]] = *col.(*interface{})
		}

		results = append(results, rowData)
	}

	jsonData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

