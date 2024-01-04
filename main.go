package main

import (
	"database/sql"
	"enigma-laundry/entity"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pass   = "7678473"
	dbname = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

func main() {
	/*
		================================== DEBUGING ONLY ==================================
	*/
	// customer := entity.Customer{Id: 4, Name: "Alhans", Phone: "081256783902", Address: "Jepara"}
	// employer := entity.Employer{Id: 4, Name: "Alhans", Phone: "089765456324", Address: "Surakarta"}
	// layanan := entity.Layanan{Id: 7, ServiceName: "Laundry Gorden", Price: 30000, Unit: "Buah"}
	// fmt.Println(customer)
	// fmt.Println(employer)
	// fmt.Println(layanan)

	// fmt.Println(connectDb())

	// customers := getAllCustomer()

	// addCustomer(customer)
	// updateCustomer(customer)
	// deleteCustomer(4)

	// addEmployer(employer)
	// updateEmployer(employer)
	// deleteEmployer(4)

	// addService(layanan)
	// updateService(layanan)
	// deleteService(7)

	arrays := getAllTransaction()
	for _, x := range arrays {
		fmt.Println(x)
	}

	/*
		================================== [END] DEBUGING ONLY ==================================
	*/
}

/*
================================== CONNECT DB FUNCTION ==================================
-> Koneksi ke Database PLSQL
*/

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected!")
	}

	return db
}

/*
================================== SCAN CUSTOMER FUNCTION ==================================
-> Mengambil data Customer dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
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
================================== SCAN EMPLOYER FUNCTION ==================================
-> Mengambil data Employer dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanEmployer(rows *sql.Rows) []entity.Employer {
	employers := []entity.Employer{}
	var err error

	for rows.Next() {
		employer := entity.Employer{}
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
================================== SCAN SERVICE FUNCTION ==================================
-> Mengambil data Layanan dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanService(rows *sql.Rows) []entity.Layanan {
	services := []entity.Layanan{}
	var err error

	for rows.Next() {
		service := entity.Layanan{}
		err := rows.Scan(&service.Id, &service.ServiceName, &service.Price, &service.Unit)

		if err != nil {
			panic(err)
		}

		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return services
}

/*
================================== SCAN TRANSACTION FUNCTION ==================================
-> Mengambil data Transaksi dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanTransaction(rows *sql.Rows) []entity.Transaksi {
	transactions := []entity.Transaksi{}
	var err error

	for rows.Next() {
		transaction := entity.Transaksi{}
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
================================== GET ALL CUSTOMER FUNCTION ==================================
-> Mengambil semua data Customer dari tabel mst_customer
-> Mengembalikan nilai berupa Array Struct of Customer
*/

func getAllCustomer() []entity.Customer {
	selectStatement := "SELECT id, name, phone, address FROM mst_customer"

	db := connectDb()
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
================================== GET ALL EMPLOYER FUNCTION ==================================
-> Mengambil semua data Employer dari tabel mst_employer
-> Mengembalikan nilai berupa Array Struct of Employer
*/

func getAllEmployer() []entity.Employer {
	selectStatement := "SELECT id, name, phone, address FROM mst_employer"

	db := connectDb()
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
================================== GET ALL SERVICE FUNCTION ==================================
-> Mengambil semua data Layanan dari tabel mst_layanan
-> Mengembalikan nilai berupa Array Struct of Layanan
*/

func getAllService() []entity.Layanan {
	selectStatement := "SELECT id, service_name, price, unit FROM mst_layanan"

	db := connectDb()
	defer db.Close()
	var err error

	rows, err := db.Query(selectStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	services := scanService(rows)
	return services
}

/*
================================== GET ALL TRANSACTION FUNCTION ==================================
-> Mengambil semua data Transaksi
-> Menggabungkan Tabel trx_bill, mst_customer dan mst_employer
*/

func getAllTransaction() []entity.Transaksi {
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

	db := connectDb()
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
================================== ADD CUSTOMER FUNCTION ==================================
-> Menambah Data Customer Baru
-> Menggunakan tabel mst_customer
*/

func addCustomer(customer entity.Customer) {
	insertStatement := "INSERT INTO mst_customer (id, name, phone, address) VALUES ($1, $2, $3, $4);"

	db := connectDb()
	defer db.Close()
	var err error

	_, err = db.Exec(insertStatement, customer.Id, customer.Name, customer.Phone, customer.Address)

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

func updateCustomer(customer entity.Customer) {
	updateStatement := "UPDATE mst_customer SET name = $2, phone=$3, address=$4 WHERE id=$1;"

	db := connectDb()
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

func deleteCustomer(id int) {
	deleteStatement := "DELETE FROM mst_customer WHERE id=$1;"

	db := connectDb()
	defer db.Close()
	var err error

	_, err = db.Exec(deleteStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

/*
================================== ADD EMPLOYER FUNCTION ==================================
-> Menambah Data Employer Baru
-> Menggunakan tabel mst_employer
*/

func addEmployer(employer entity.Employer) {
	insertStatement := "INSERT INTO mst_employer (id, name, phone, address) VALUES ($1, $2, $3, $4);"

	db := connectDb()
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

func updateEmployer(employer entity.Employer) {
	updateStatement := "UPDATE mst_employer SET name = $2, phone=$3, address=$4 WHERE id=$1;"

	db := connectDb()
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

func deleteEmployer(id int) {
	deleteStatement := "DELETE FROM mst_employer WHERE id=$1;"

	db := connectDb()
	defer db.Close()
	var err error

	_, err = db.Exec(deleteStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

/*
================================== ADD SERVICE FUNCTION ==================================
-> Menambah Data Layanan Baru
-> Menggunakan tabel mst_layanan
*/

func addService(service entity.Layanan) {
	insertStatement := "INSERT INTO mst_layanan (id, service_name, price, unit) VALUES ($1, $2, $3, $4);"

	db := connectDb()
	defer db.Close()
	var err error

	_, err = db.Exec(insertStatement, service.Id, service.ServiceName, service.Price, service.Unit)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

/*
================================== UPDATE SERVICE FUNCTION ==================================
-> Mengubah Data Layanan Berdasarkan Id
-> Menggunakan tabel mst_layanan
*/

func updateService(service entity.Layanan) {
	updateStatement := "UPDATE mst_layanan SET service_name = $2, price=$3, unit=$4 WHERE id=$1;"

	db := connectDb()
	defer db.Close()
	var err error

	_, err = db.Exec(updateStatement, service.Id, service.ServiceName, service.Price, service.Unit)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

/*
================================== DELETE SERVICE FUNCTION ==================================
-> Menghapus Data Layanan Berdasarkan Id
-> Menggunakan tabel mst_layanan
*/

func deleteService(id int) {
	deleteStatement := "DELETE FROM mst_layanan WHERE id=$1;"

	db := connectDb()
	defer db.Close()
	var err error

	_, err = db.Exec(deleteStatement, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}
