package entity

import (
	"database/sql"
	"enigma-laundry/db"
	"fmt"
)

type Employer struct {
	Id      int
	Name    string
	Phone   string
	Address string
}

/*
================================== SCAN EMPLOYER FUNCTION ==================================
-> Mengambil data Employer dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanEmployer(rows *sql.Rows) []Employer {
	employers := []Employer{}
	var err error

	for rows.Next() {
		employer := Employer{}
		err := rows.Scan(&employer.Id, &employer.Name, &employer.Phone, &employer.Address)

		if err != nil {
			panic(err)
		}

		employers = append(employers, employer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return employers
}

/*
================================== GET ALL EMPLOYER FUNCTION ==================================
-> Mengambil semua data Employer dari tabel mst_employer
-> Mengembalikan nilai berupa Array Struct of Employer
*/

func GetAllEmployer() []Employer {
	selectStatement := "SELECT id, name, phone, address FROM mst_employer"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	rows, err := db.Query(selectStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	employers := scanEmployer(rows)
	return employers
}

/*
================================== ADD EMPLOYER FUNCTION ==================================
-> Menambah Data Employer Baru
-> Menggunakan tabel mst_employer
*/

func AddEmployer(employer Employer) {
	insertStatement := "INSERT INTO mst_employer (id, name, phone, address) VALUES ($1, $2, $3, $4);"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	_, err = db.Exec(insertStatement, employer.Id, employer.Name, employer.Phone, employer.Address)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

/*
================================== UPDATE EMPLOYER FUNCTION ==================================
-> Mengubah Data Employer Berdasarkan Id
-> Menggunakan tabel mst_employer
*/

func UpdateEmployer(employer Employer) {
	updateStatement := "UPDATE mst_employer SET name = $2, phone=$3, address=$4 WHERE id=$1;"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	_, err = db.Exec(updateStatement, employer.Id, employer.Name, employer.Phone, employer.Address)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

/*
================================== DELETE EMPLOYER FUNCTION ==================================
-> Menghapus Data Employer Berdasarkan Id
-> Menggunakan tabel mst_employer
*/

func DeleteEmployer(id int) {
	deleteStatement := "DELETE FROM mst_employer WHERE id=$1;"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	_, err = db.Exec(deleteStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}
