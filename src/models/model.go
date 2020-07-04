package models

import (
	"database/sql"

	"github.com/htoomgt/go-gin-rest-api-tutorial/src/configs"
)

func getDbConnection() (dbC *sql.DB, err error) {
	var username = configs.MySQLDbUsername
	var password = configs.MySQLDbPassword
	var host = configs.MySQLDbHost
	var port = configs.MySQLDbPort
	var database = configs.MySQLDbNAME

	dbC, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?parseTime=true")
	if err != nil {
		return
	}
	return
}
