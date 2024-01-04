package entity

import (
	"database/sql"
	"enigma-laundry/db"
	"fmt"
)

type Transaksi struct {
	Id            int
	NoTransaction string
	DateIn        string
	DateOut       string
	CustomerId    int
	CustomerName  string
	CustomerPhone string
	EmployerId    int
	EmployerName  string
}

type DetailTransaksi struct {
	Id          int
	ServiceId   int
	ServiceName string
	Price       int
	Unit        string
	Quantity    int
	TotalPrice  int
}

/*
================================== SCAN TRANSACTION FUNCTION ==================================
-> Mengambil data Transaksi dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanTransaction(rows *sql.Rows) []Transaksi {
	transactions := []Transaksi{}
	var err error

	for rows.Next() {
		transaction := Transaksi{}
		err := rows.Scan(&transaction.Id, &transaction.DateIn, &transaction.DateOut, &transaction.NoTransaction, &transaction.CustomerId, &transaction.CustomerName, &transaction.CustomerPhone, &transaction.EmployerId, &transaction.EmployerName)

		if err != nil {
			panic(err)
		}

		transactions = append(transactions, transaction)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return transactions
}

/*
================================== SCAN DETAIL TRANSACTION FUNCTION ==================================
-> Mengambil data Detail Transaksi dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanDetailTransaction(rows *sql.Rows) []DetailTransaksi {
	arrays := []DetailTransaksi{}
	var err error

	for rows.Next() {
		detailTrx := DetailTransaksi{}
		err := rows.Scan(&detailTrx.Id, &detailTrx.ServiceId, &detailTrx.ServiceName, &detailTrx.Price, &detailTrx.Unit, &detailTrx.Quantity, &detailTrx.TotalPrice)

		if err != nil {
			panic(err)
		}

		arrays = append(arrays, detailTrx)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return arrays
}

/*
================================== GET ALL TRANSACTION FUNCTION ==================================
-> Mengambil semua data Transaksi
-> Menggabungkan Tabel trx_bill, mst_customer dan mst_employer
*/

func GetAllTransaction() []Transaksi {
	selectStatement := `SELECT  
t.id,
t.date_in,
t.date_out,
t.no_trx,
t.customer_id,
c.name as customer_name,
c.phone as customer_phone,
t.employer_id,
e.name as employer_name
FROM trx_bill as t
JOIN mst_customer as c ON t.customer_id = c.id
JOIN mst_employer as e ON t.employer_id = e.id;`

	db := db.ConnectDB()
	defer db.Close()
	var err error

	rows, err := db.Query(selectStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	transactions := scanTransaction(rows)
	return transactions
}

/*
================================== GET DETAIL TRANSACTION FUNCTION ==================================
-> Mengambil detail data Transaksi tertentu berdasarkan Id Transaksi
-> Menggabungkan Tabel trx_bill_detail dan mst_layanan
*/

func GetDetailTransaction(id int) []DetailTransaksi {
	selectStatement := `SELECT
td.trx_bill_id as transaction_id,
td.service_id,
l.service_name,
l.price,
l.unit,
td.quantity,
SUM(l.price*td.quantity) as total
FROM trx_bill_detail as td
JOIN mst_layanan as l ON td.service_id = l.id
WHERE td.trx_bill_id = $1
GROUP BY td.trx_bill_id, td.service_id, l.service_name, l.price, l.unit, td.quantity;`

	db := db.ConnectDB()
	defer db.Close()
	var err error

	rows, err := db.Query(selectStatement, id)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	detailTrx := scanDetailTransaction(rows)
	return detailTrx
}

/*
================================== MAKE TRANSACTION FUNCTION ==================================
-> Insert data Transaksi
-> Caranya dengan Insert data di tabel trx_bill kemudian insert data di tabel trx_bill_detail
*/

func MakeTransaction(transaction Transaksi, detailTrx []DetailTransaksi) {
	db := db.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	insertTransaction(transaction, tx)
	insertDetailTransaction(detailTrx, transaction, tx)

	err = tx.Commit()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaction Commited!")
	}
}

func insertTransaction(transaction Transaksi, tx *sql.Tx) {
	insertStatement := "INSERT INTO trx_bill (id, no_trx, date_in, date_out, customer_id, employer_id) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := tx.Exec(insertStatement, transaction.Id, transaction.NoTransaction, transaction.DateIn, transaction.DateOut, transaction.CustomerId, transaction.EmployerId)

	validate(err, "Insert Transaction", tx)
}

func insertDetailTransaction(detailTrx []DetailTransaksi, transaction Transaksi, tx *sql.Tx) {
	insertStatement := "INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES ($1, $2, $3, $4);"

	for _, trx := range detailTrx {
		_, err := tx.Exec(insertStatement, trx.Id, trx.ServiceId, trx.Quantity, transaction.Id)

		validate(err, "Insert Detail Transaction", tx)
	}
	fmt.Println("Successfully Insert All Detail Transaction")
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback!")
	} else {
		fmt.Println("Successfully " + message + " Data!")
	}
}
