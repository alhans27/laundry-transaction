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
	// customer := entity.Customer{Id: 1, Name: "Jessica", Phone: "089765456324", Address: "Jakarta"}
	// employer := entity.Employer{Id: 1, Name: "Mirna", Phone: "089765456324", Address: "Surabaya"}
	// layanan := entity.Layanan{Id: 1, ServiceName: "Cuci", Price: 5000, Unit: "KG"}
	// fmt.Println(customer)
	// fmt.Println(employer)
	// fmt.Println(layanan)

	// fmt.Println(connectDb())

	// customers := getAllCustomer()
	employers := getAllService()

	for _, emp := range employers {
		fmt.Println(emp)
	}
}

/*
== CONNECT DB FUNCTION ==
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
== SCAN CUSTOMER FUNCTION ==
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
== SCAN EMPLOYER FUNCTION ==
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
== SCAN SERVICE FUNCTION ==
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
== GET ALL CUSTOMER FUNCTION ==
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
== GET ALL EMPLOYER FUNCTION ==
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
== GET ALL SERVICE FUNCTION ==
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
