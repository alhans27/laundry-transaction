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
	customer := entity.Customer{Id: 1, Name: "Jessica", Phone: "089765456324", Address: "Jakarta"}
	employer := entity.Employer{Id: 1, Name: "Mirna", Phone: "089765456324", Address: "Surabaya"}
	layanan := entity.Layanan{Id: 1, ServiceName: "Cuci", Price: 5000, Unit: "KG"}
	fmt.Println(customer)
	fmt.Println(employer)
	fmt.Println(layanan)

	fmt.Println(connectDb())
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()

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
