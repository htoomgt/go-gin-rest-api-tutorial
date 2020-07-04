package models

import "database/sql"

func getDbConnection() (dbC *sql.DB, err error) {
	dbC, err = sql.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/db_test?parseTime=true")
	if err != nil {
		return
	}
	return
}
