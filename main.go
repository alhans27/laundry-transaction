package main

import (
	// "bufio"
	"enigma-laundry/entity"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	isRun    = true
	t        = table.NewWriter()
	customer = entity.Customer{}
	employer = entity.Employer{}
	layanan  = entity.Layanan{}
)

// var scanner = bufio.NewScanner(os.Stdin)

func main() {
	t.SetOutputMirror(os.Stdout)
	/*
		================================== DEBUGING ONLY ==================================
	*/
	// customer := entity.Customer{Id: 4, Name: "Alhans", Phone: "081256783902", Address: "Surakarta"}
	// employer := entity.Employer{Id: 4, Name: "Alhans", Phone: "089765456324", Address: "Surakarta"}
	// layanan := entity.Layanan{Id: 7, ServiceName: "Laundry Gorden", Price: 30000, Unit: "Buah"}
	// fmt.Println(customer)
	// fmt.Println(employer)
	// fmt.Println(layanan)

	// fmt.Println(connectDb())

	// entity.AddCustomer(customer)
	// entity.UpdateCustomer(customer)
	// entity.DeleteCustomer(4)

	// customers := entity.GetAllCustomer()
	// for _, x := range customers {
	// 	fmt.Println(x)
	// }

	// addEmployer(employer)
	// updateEmployer(employer)
	// deleteEmployer(4)

	// addService(layanan)
	// updateService(layanan)
	// deleteService(7)

	// arrays1 := getAllTransaction()
	// for _, x := range arrays1 {
	// 	fmt.Println(x)
	// 	arrays2 := getDetailTransaction(x.Id)
	// 	for _, y := range arrays2 {
	// 		fmt.Println(y)
	// 	}
	// 	fmt.Println("========================================")
	// }

	// transaction := entity.Transaksi{Id: 3, NoTransaction: "23456", DateIn: "2024-01-04", DateOut: "2024-01-07", CustomerId: 2, EmployerId: 2}

	// var arrays = []entity.DetailTransaksi{
	// 	{Id: 7, ServiceId: 3, Quantity: 2},
	// 	{Id: 8, ServiceId: 1, Quantity: 3},
	// 	{Id: 9, ServiceId: 2, Quantity: 1},
	// 	{Id: 10, ServiceId: 6, Quantity: 4},
	// }
	// arrays = append(arrays, detailTrx)

	// makeTransaction(transaction, arrays)

	/*
		================================== [END] DEBUGING ONLY ==================================
	*/

	for isRun {
		choice := displayMainMenu()
		switch choice {
		case "1":
			choice = displayCustomerMenu()
			switch choice {
			case "1":
				displayAllCustomer()
			case "2":
				displayAddCustomer()
			}
		case "2":
			choice = displayEmployerMenu()
			switch choice {
			case "1":
				displayAllEmployer()
			}
		case "3":
			choice = displayServiceMenu()
			switch choice {
			case "1":
				displayAllService()
			}
		case "4":
			choice = displayTransactionMenu()
			switch choice {
			case "1":
				displayAllTransaction()
			}
		case "5":
			isRun = false
		}
	}
}

func displayMainMenu() string {
	var choice string
	fmt.Println("======================================================================")
	fmt.Println("====                        ENIGMA LAUNDRY                        ====")
	fmt.Println("======================================================================")
	fmt.Println("|||      MAIN MENU      |||")
	fmt.Println("---------------------------")
	fmt.Println("1. Data Customer")
	fmt.Println("2. Data Karyawan")
	fmt.Println("3. Data Layanan")
	fmt.Println("4. Data Transaksi")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------------------")
	fmt.Print("Menu yang Anda Pilih? ")
	fmt.Scan(&choice)
	fmt.Println()

	return choice
}

func displayCustomerMenu() string {
	var choice string
	fmt.Println("||      CUSTOMER MENU      ||")
	fmt.Println("---------------------------")
	fmt.Println("1. Lihat Data Customer")
	fmt.Println("2. Tambah Data Customer")
	fmt.Println("3. Update Data Customer")
	fmt.Println("4. Delete Data Customer")
	fmt.Println("5. Back to Main Menu")
	fmt.Println("-----------------------------------")
	fmt.Print("Menu yang Anda Pilih? ")
	fmt.Scan(&choice)
	fmt.Println()

	return choice
}

func displayEmployerMenu() string {
	var choice string
	fmt.Println("||      EMPLOYER MENU      ||")
	fmt.Println("---------------------------")
	fmt.Println("1. Lihat Data Karyawan")
	fmt.Println("2. Tambah Data Karyawan")
	fmt.Println("3. Update Data Karyawan")
	fmt.Println("4. Delete Data Karyawan")
	fmt.Println("5. Back to Main Menu")
	fmt.Println("-----------------------------------")
	fmt.Print("Menu yang Anda Pilih? ")
	fmt.Scan(&choice)
	fmt.Println()

	return choice
}

func displayServiceMenu() string {
	var choice string
	fmt.Println("||      SERVICE MENU      ||")
	fmt.Println("---------------------------")
	fmt.Println("1. Lihat Data Layanan")
	fmt.Println("2. Tambah Data Layanan")
	fmt.Println("3. Update Data Layanan")
	fmt.Println("4. Delete Data Layanan")
	fmt.Println("5. Back to Main Menu")
	fmt.Println("-----------------------------------")
	fmt.Print("Menu yang Anda Pilih? ")
	fmt.Scan(&choice)
	fmt.Println()

	return choice
}

func displayTransactionMenu() string {
	var choice string
	fmt.Println("||      TRANSACTION MENU      ||")
	fmt.Println("---------------------------")
	fmt.Println("1. Lihat Semua Data Transaksi")
	fmt.Println("2. Tambah Data Transaksi")
	fmt.Println("3. Back to Main Menu")
	fmt.Println("-----------------------------------")
	fmt.Print("Menu yang Anda Pilih? ")
	fmt.Scan(&choice)
	fmt.Println()

	return choice
}

func displayAllCustomer() {
	var choice string
	customers := entity.GetAllCustomer()
	fmt.Println()
	fmt.Println("|\t\t\tCUSTOMER ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Nama Customer", "Customer Id", "Nomor Handphone", "Alamat"})
	for i, cust := range customers {
		t.AppendRow([]interface{}{i + 1, cust.Name, cust.Id, cust.Phone, cust.Address})
	}
	t.Render()

	fmt.Println()

	for choice != "Y" && choice != "y" || choice != "N" && choice != "n" {
		fmt.Print("Back to Main Menu? (Y/n)\n")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {
			isRun = true
			break
		} else if choice == "N" || choice == "n" {
			isRun = false
			break
		} else {
			fmt.Println("Anda menginputkan pilihan yang salah. Coba lagi!")
		}
	}
}

func displayAllEmployer() {
	var choice string
	employers := entity.GetAllEmployer()
	fmt.Println()
	fmt.Println("|\t\t\tKARYAWAN ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Nama Karyawan", "Karyawan Id", "Nomor Handphone", "Alamat"})
	for i, emp := range employers {
		t.AppendRow([]interface{}{i + 1, emp.Name, emp.Id, emp.Phone, emp.Address})
	}
	t.Render()

	fmt.Println()

	for choice != "Y" && choice != "y" || choice != "N" && choice != "n" {
		fmt.Print("Back to Main Menu? (Y/n)\n")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {
			isRun = true
			break
		} else if choice == "N" || choice == "n" {
			isRun = false
			break
		} else {
			fmt.Println("Anda menginputkan pilihan yang salah. Coba lagi!")
		}
	}
}

func displayAllService() {
	var choice string
	services := entity.GetAllService()
	fmt.Println()
	fmt.Println("|\t\t\tLAYANAN ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Nama Layanan", "Harga", "Satuan"})
	for i, serv := range services {
		t.AppendRow([]interface{}{i + 1, serv.ServiceName, serv.Price, serv.Unit})
	}
	t.Render()

	fmt.Println()

	for choice != "Y" && choice != "y" || choice != "N" && choice != "n" {
		fmt.Print("Back to Main Menu? (Y/n)\n")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {
			isRun = true
			break
		} else if choice == "N" || choice == "n" {
			isRun = false
			break
		} else {
			fmt.Println("Anda menginputkan pilihan yang salah. Coba lagi!")
		}
	}
}

func displayAllTransaction() {
	var choice string
	transactions := entity.GetAllTransaction()
	fmt.Println()
	fmt.Println("|\t\t\tTRANSAKSI ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Tanggal Masuk", "Tanggal Keluar", "Nomor Transaksi", "Diterima Oleh", "Nama Customer", "No Hp Customer"})
	for i, trx := range transactions {
		dateIn, err := time.Parse("2006-01-02T15:04:05Z", trx.DateIn)
		dateOut, err := time.Parse("2006-01-02T15:04:05Z", trx.DateOut)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		trx.DateIn = dateIn.Format("2006-01-02")
		trx.DateOut = dateOut.Format("2006-01-02")
		t.AppendRow([]interface{}{i + 1, trx.DateIn, trx.DateOut, trx.NoTransaction, trx.EmployerName, trx.CustomerName, trx.CustomerPhone})
	}
	t.Render()

	fmt.Println()

	for choice != "Y" && choice != "y" || choice != "N" && choice != "n" {
		fmt.Print("Ingin Melihat Detail Transaksi? (Y/n)\n")
		fmt.Scan(&choice)
		if choice == "Y" || choice == "y" {
			isRun = true
			displayDetailTransaction()
			break
		} else if choice == "N" || choice == "n" {
			isRun = true
			break
		} else {
			fmt.Println("Anda menginputkan pilihan yang salah. Coba lagi!")
		}
	}
}

func displayDetailTransaction() {
	var choice string
	var totalHarga int

	fmt.Print("Masukkan Id Transaksi? ")
	fmt.Scan(&choice)
	id, err := strconv.Atoi(choice)

	if err != nil {
		panic(err)
	}

	arrays := entity.GetDetailTransaction(id)

	fmt.Println()
	fmt.Println("|\t\t\tDETAIL TRANSAKSI ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Nama Layanan", "Harga", "Satuan", "Jumlah", "Total"})

	for i, trx := range arrays {
		totalHarga += trx.TotalPrice
		t.AppendRow([]interface{}{i + 1, trx.ServiceName, trx.Price, trx.Unit, trx.Quantity, trx.TotalPrice})
	}

	t.AppendFooter(table.Row{"", "", "", "", "Total", totalHarga})

	t.Render()
	fmt.Println()
}

func displayAddCustomer() {
	var choice string
	isLooping := true
	fmt.Println("Tambah Customer Baru")
	fmt.Println("===========================")
	for isLooping {
		fmt.Print("Masukkan Nama Customer : ")
		fmt.Scan(&customer.Name)
		if customer.Name == "" || customer.Name == " " {
			fmt.Println("Nama Tidak Boleh Kosong")
		} else if len(customer.Name) < 2 {
			fmt.Println("Nama Tidak Boleh Kosong")
		}
		fmt.Print("Masukkan Nomor Handphone Customer : ")
		fmt.Scan(&customer.Phone)
		fmt.Print("Masukkan Alamat Customer : ")
		fmt.Scan(&customer.Address)
		fmt.Println("--------------------------------------------")
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Nama Customer : %s\nNomor Hp : %s\nAlamat : %s\n", customer.Name, customer.Phone, customer.Address)
		fmt.Print("Yakin Data Sudah Benar? (Y/n) ")
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			entity.AddCustomer(customer)
			break
		}
	}
}

func resetTable() {
	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()
}
