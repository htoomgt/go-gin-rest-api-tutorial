package models

import (
	_ "github.com/go-sql-driver/mysql" //this package is loaded for database tasks
)

/*Person type strucutre for processing database and serialization*/
type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Gender    string `json:"gender" form:"gender"`
	Age       int    `json:"age" form:"age"`
	Confirmed bool   `json:"confirmed" form:"confirmed"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}

/*GetByID function retrives record by person id*/
func (p Person) GetByID() (person Person, err error) {
	db, err := getDbConnection()
	row := db.QueryRow("SELECT id, first_name, last_name, email, gender, age, created_at, updated_at FROM tbl_persons WHERE id=?", p.ID)
	err = row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.Gender, &person.Age, &person.CreatedAt, &person.UpdatedAt)
	if err != nil {
		return
	}

	return
}

/*GetAll function retrieves all person records from database*/
func (p Person) GetAll() (persons []Person, err error) {
	db, err := getDbConnection()
	if err != nil {
		return
	}
	var sqlStmt string = "SELECT id, first_name, last_name, email, gender, age, created_at, updated_at FROM tbl_persons"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.Gender, &person.Age, &person.CreatedAt, &person.UpdatedAt)
		persons = append(persons, person)
	}
	defer rows.Close()
	return

}

/*add()Person to insert a person to database table*/
func (p Person) add() (ID int, err error) {
	db, err := getDbConnection()
	if err != nil {
		return
	}

	stmt, err := db.Prepare("INSERT INTO tbl_psersons (first_name, last_name) VALUES(?, ?)")
	if err != nil {
		return
	}

	rs, err := stmt.Exec(p.FirstName, p.LastName)
	if err != nil {
		return
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return
	}

	ID = int(id)
	defer stmt.Close()
	return

}
