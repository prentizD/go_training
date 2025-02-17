package models

import (
	"database/sql"
	"time"
)

type Logging struct {
	Logging_id int
	Route      string
	Date       time.Time
}

type LoggingModel struct {
	DB *sql.DB
}

func (m *LoggingModel) Insert(route string) (int, error) {

	stmt := `INSERT INTO logging (route, date) VALUES (?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, route)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
