package entity

import (
	"database/sql"
	"enigma-laundry/db"
	"fmt"
)

type Customer struct {
	Id      int
	Name    string
	Phone   string
	Address string
}

/*
================================== SCAN CUSTOMER FUNCTION ==================================
-> Mengambil data Customer dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanCustomer(rows *sql.Rows) []Customer {
	customers := []Customer{}
	var err error

	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address)

		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return customers
}

func GetLastIdCustomer() int {
	selectStatement := "SELECT id FROM mst_customer ORDER BY id DESC LIMIT 1"
	var id int

	db := db.ConnectDB()
	defer db.Close()
	var err error

	err = db.QueryRow(selectStatement).Scan(&id)

	if err != nil {
		panic(err)
	}
	return id
}

/*
================================== GET ALL CUSTOMER FUNCTION ==================================
-> Mengambil semua data Customer dari tabel mst_customer
-> Mengembalikan nilai berupa Array Struct of Customer
*/

func GetAllCustomer() []Customer {
	selectStatement := "SELECT id, name, phone, address FROM mst_customer"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	rows, err := db.Query(selectStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	customers := scanCustomer(rows)
	return customers
}

/*
================================== ADD CUSTOMER FUNCTION ==================================
-> Menambah Data Customer Baru
-> Menggunakan tabel mst_customer
*/

func AddCustomer(customer Customer) {
	id := GetLastIdCustomer()
	insertStatement := "INSERT INTO mst_customer (id, name, phone, address) VALUES ($1, $2, $3, $4);"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	_, err = db.Exec(insertStatement, id, customer.Name, customer.Phone, customer.Address)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

/*
================================== UPDATE CUSTOMER FUNCTION ==================================
-> Mengubah Data Customer Berdasarkan Id
-> Menggunakan tabel mst_customer
*/

func UpdateCustomer(customer Customer) {
	updateStatement := "UPDATE mst_customer SET name = $2, phone=$3, address=$4 WHERE id=$1;"

	db := db.ConnectDB()
	defer db.Close()
	var err error

	_, err = db.Exec(updateStatement, customer.Id, customer.Name, customer.Phone, customer.Address)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

/*
================================== DELETE CUSTOMER FUNCTION ==================================
-> Menghapus Data Customer Berdasarkan Id
-> Menggunakan tabel mst_customer
*/

func DeleteCustomer(id int) {
	deleteStatement := "DELETE FROM mst_customer WHERE id=$1;"

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
