package entity

import (
	"database/sql"
	"enigma-laundry/db"
	"fmt"
)

type Layanan struct {
	Id          int
	ServiceName string
	Price       int
	Unit        string
}

/*
================================== SCAN SERVICE FUNCTION ==================================
-> Mengambil data Layanan dari hasil db.Query() dan memasukkannya ke dalam Array Struct
*/

func scanService(rows *sql.Rows) []Layanan {
	services := []Layanan{}
	var err error

	for rows.Next() {
		service := Layanan{}
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
================================== GET ALL SERVICE FUNCTION ==================================
-> Mengambil semua data Layanan dari tabel mst_layanan
-> Mengembalikan nilai berupa Array Struct of Layanan
*/

func GetAllService() []Layanan {
	selectStatement := "SELECT id, service_name, price, unit FROM mst_layanan"

	db := db.ConnectDB()
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
================================== ADD SERVICE FUNCTION ==================================
-> Menambah Data Layanan Baru
-> Menggunakan tabel mst_layanan
*/

func AddService(service Layanan) {
	insertStatement := "INSERT INTO mst_layanan (id, service_name, price, unit) VALUES ($1, $2, $3, $4);"

	db := db.ConnectDB()
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

func UpdateService(service Layanan) {
	updateStatement := "UPDATE mst_layanan SET service_name = $2, price=$3, unit=$4 WHERE id=$1;"

	db := db.ConnectDB()
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

func DeleteService(id int) {
	deleteStatement := "DELETE FROM mst_layanan WHERE id=$1;"

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