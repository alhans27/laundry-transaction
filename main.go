package main

import (
	"enigma-laundry/entity"
	"fmt"
)

func main() {
	customer := entity.Customer{Id: 1, Name: "Jessica", Phone: "089765456324", Address: "Jakarta"}
	employer := entity.Employer{Id: 1, Name: "Mirna", Phone: "089765456324", Address: "Surabaya"}
	layanan := entity.Layanan{Id: 1, ServiceName: "Cuci", Price: 5000, Unit: "KG"}
	fmt.Println(customer)
	fmt.Println(employer)
	fmt.Println(layanan)
}
