package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Open a new database connection
	//db, err := sql.Open("mysql", "username:password@tcp(IP address:Port)/dbname")
	db, err := sql.Open("mysql", "root:IST888IST888@tcp(127.0.0.1:3306)/School")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// CREATE operation
	_, err = db.Exec("INSERT INTO tblStudents(fldStudentId, fldFName, fldLName, fldGPA, fldCurrentCredits, fldTotalCredits) "+
		"VALUES (?, ?, ?, ?, ?, ?)", 1001, "Joe", "Oakes", 4.0, 12, 100)
	if err != nil {
		log.Fatal(err)
	}

	// READ operation
	rows, err := db.Query("SELECT fldStudentId, fldFName, fldLName, fldGPA, fldCurrentCredits, fldTotalCredits FROM tblStudents")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var fldStudentId int
		var fldFName string
		var fldLName string
		var fldGPA float32
		var fldCurrentCredits int
		var fldTotalCredits int
		err := rows.Scan(&fldStudentId, &fldFName, &fldLName, &fldGPA, &fldCurrentCredits, &fldTotalCredits)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fldStudentId, fldFName, fldLName, fldGPA, fldCurrentCredits, fldTotalCredits)
	}

	// UPDATE operation
	_, err = db.Exec("UPDATE tblStudents SET fldGPA = ? WHERE fldStudentId = ?", 1000, 4.0)
	if err != nil {
		log.Fatal(err)
	}

	// DELETE operation
	/*
		_, err = db.Exec("DELETE FROM tblStudents WHERE fldStudentId = ?", 1000)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
