package entity

import (
	"database/sql"
	"enigma-laundry/db"
	"fmt"
)

type Customer struct {
	Id      int    `validate:"numeric"`
	Name    string `validate:"required,alpha,min=3,max=50"`
	Phone   string `validate:"required,min=11,max=13"`
	Address string `validate:"required"`
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

/*
================================== GET LAST ID CUSTOMER FUNCTION ==================================
-> Mengambil ID Terakhir dari data Customer pada tabel mst_customer
-> Mengembalikan nilai berupa int
*/

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
	return id + 1
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
	selectStatement := "SELECT customer_id FROM trx_bill WHERE customer_id = $1;"
	deleteStatement := "DELETE FROM mst_customer WHERE id=$1;"

	var custId int

	db := db.ConnectDB()
	defer db.Close()

	err := db.QueryRow(selectStatement, id).Scan(&custId)
	if custId != 0 && err == nil {
		fmt.Println("[MESSAGE] Maaf ID Tersebut tidak dapat DIHAPUS karena telah digunakan dalam Transaksi")
	} else if custId == 0 && err != nil {
		_, err = db.Exec(deleteStatement, id)
		fmt.Println("[MESSAGE] Successfully Delete Data!")
	} else if err != nil {
		panic(err)
	}
}
