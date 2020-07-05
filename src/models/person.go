package models

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //this package is loaded for database tasks
)

/*Person type strucutre for processing database and serialization*/
type Person struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name" form:"first_name"`
	LastName  string `db:"last_name" json:"last_name" form:"last_name"`
	Email     string `db:"email" json:"email" form:"email"`
	Password  string `db:"password" json:"password" form:"password"`
	Gender    string `db:"gender" json:"gender" form:"gender"`
	Age       int    `db:"age" json:"age" form:"age"`
	Confirmed bool   `db:"confirmed" json:"confirmed" form:"confirmed"`
	CreatedAt string `db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at" form:"updated_at"`
}

/*GetByID function retrives record by person id*/
func (p *Person) GetByID() (person Person, err error) {
	db, err := getDbConnection()
	row := db.QueryRow("SELECT id, first_name, last_name, email, gender, age, created_at, updated_at FROM tbl_persons WHERE id=?", p.ID)
	err = row.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.Gender, &person.Age, &person.CreatedAt, &person.UpdatedAt)
	if err != nil {
		return
	}

	return
}

/*GetAll function retrieves all person records from database*/
func (p *Person) GetAll() (persons []Person, err error) {
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

/*Add is a function to insert a person to database table*/
func (p *Person) Add() (ID int, err error) {
	db, err := getDbConnection()
	dt := time.Now()
	nowDt := dt.Format("2006-01-02 15:04:05")
	if err != nil {
		return
	}

	stmt, err := db.Prepare("INSERT INTO tbl_persons (first_name, last_name, email, password, gender, age, confirmed, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}

	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Email, p.Password, p.Gender, p.Age, false, nowDt)
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

/*UpdateAllByID is a function to update person's all fields by ID at database table*/
func (p *Person) UpdateAllByID() (ra int64, err error) {
	db, err := getDbConnection()
	var sqlStmt string = `UPDATE tbl_persons SET 
		first_name=?, last_name=?, 
		email=?, password=?, 
		gender=?, age=?, confirmed=? 
	WHERE ID=? `
	/* sqlStmt += "first_name=?, "
	sqlStmt += "last_name=?, "
	sqlStmt += "email=?, "
	sqlStmt += "password=?,"
	sqlStmt += "gender=?, "
	sqlStmt += "age=?, "
	sqlStmt += "confirmed=? "
	sqlStmt += "WHERE ID=? " */

	stmt, err := db.Prepare(sqlStmt)
	defer stmt.Close()
	if err != nil {
		return
	}

	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Email, p.Password, p.Gender, p.Age, p.Confirmed, p.ID)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return

}

/*DeleteByID is a function to delete person's data row at database by ID*/
func (p *Person) DeleteByID() (ra int64, err error) {
	db, err := getDbConnection()
	rs, err := db.Exec("DELETE FROM tbl_persons WHERE id=?", p.ID)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err = rs.RowsAffected()
	return
}
