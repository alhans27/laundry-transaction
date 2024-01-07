package main

import (
	"bufio"
	"enigma-laundry/entity"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jedib0t/go-pretty/v6/table"
)

var (
	isRun            = true
	t                = table.NewWriter()
	customer         = entity.Customer{}
	employer         = entity.Employer{}
	layanan          = entity.Layanan{}
	transaksi        = entity.Transaksi{}
	detail_transaksi = entity.DetailTransaksi{}
	validate         = validator.New()
	scanner          = bufio.NewScanner(os.Stdin)
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
	// entity.DeleteService(5)

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

	/*
		================================== MAIN CODE ==================================
		-> Menampilkan Menu-menu yang tersedia
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
			case "3":
				displayUpdateCustomer()
			case "4":
				displayDeleteCustomer()
			}
		case "2":
			choice = displayEmployerMenu()
			switch choice {
			case "1":
				displayAllEmployer()
			case "2":
				displayAddEmployer()
			case "3":
				displayUpdateEmployer()
			case "4":
				displayDeleteEmployer()
			}
		case "3":
			choice = displayServiceMenu()
			switch choice {
			case "1":
				displayAllService()
			case "2":
				displayAddService()
			case "3":
				displayUpdateService()
			case "4":
				displayDeleteService()
			}
		case "4":
			choice = displayTransactionMenu()
			switch choice {
			case "1":
				displayAllTransaction()
			case "2":
				displayAddTransaction()
			}
		case "5":
			isRun = false
		}
	}
}

/*
================================== DISPLAY MAIN MENU FUNCTION ==================================
-> Menampilkan Pilihan Menu Utama Dari Aplikasi
*/

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

/*
================================== DISPLAY CUSTOMER MENU FUNCTION ==================================
-> Menampilkan Pilihan Menu Tentang Customer
*/

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

/*
================================== DISPLAY EMPLOYER MENU FUNCTION ==================================
-> Menampilkan Pilihan Menu Tentang Karyawan
*/

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

/*
================================== DISPLAY SERVICE MENU FUNCTION ==================================
-> Menampilkan Pilihan Menu Tentang Layanan
*/

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

/*
================================== DISPLAY TRANSACTION MENU FUNCTION ==================================
-> Menampilkan Pilihan Menu Tentang Transaksi
*/

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

/*
================================== DISPLAY ALL CUSTOMER FUNCTION ==================================
-> UI Untuk Menampilkan Semua Data Customer
*/

func displayAllCustomer() {
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
}

/*
================================== DISPLAY ADD CUSTOMER FUNCTION ==================================
-> UI Untuk Menambahkan Data Customer
*/

func displayAddCustomer() {
	var choice string
	isLooping := true
	fmt.Println("Tambah Customer Baru")
	fmt.Println("===========================")
	for isLooping {
		fmt.Print("Masukkan Nama Customer : ")
		fmt.Scan(&customer.Name)
		fmt.Print("Masukkan Nomor Handphone Customer : ")
		fmt.Scan(&customer.Phone)
		fmt.Print("Masukkan Alamat Customer : ")
		fmt.Scan(&customer.Address)
		fmt.Println("--------------------------------------------")
		err := validate.Struct(customer)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldErr := range validationErrors {
				if fieldErr.Field() == "Name" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "alpha":
						message = "Hanya boleh berupa huruf!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nama Customer", message)
				}
				if fieldErr.Field() == "Phone" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nomor Handphone Customer", message)
				}
				if fieldErr.Field() == "Address" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Alamat Customer", message)
				}
			}
			fmt.Println("--------------------------------------------")
			continue
		}
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Nama Customer : %s\nNomor Hp : %s\nAlamat : %s\n", customer.Name, customer.Phone, customer.Address)

		for true {
			fmt.Print("Apakah Data Tersebut Sudah Benar? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
				entity.AddCustomer(customer)
				isLooping = false
				break
			} else if choice == "n" {
				isLooping = true
				break
			} else {
				fmt.Println("Pilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
}

/*
================================== DISPLAY UPDATE CUSTOMER FUNCTION ==================================
-> UI Untuk Mengupdate Data Customer
*/

func displayUpdateCustomer() {
	var isLooping = true
	var choice string

	fmt.Println("Update Data Customer")
	fmt.Println("===========================")
	displayAllCustomer()
	fmt.Print("Masukkan Id Customer yang hendak di Update : ")
	fmt.Scan(&customer.Id)

	for isLooping {
		fmt.Print("Masukkan Nama Customer : ")
		fmt.Scan(&customer.Name)
		fmt.Print("Masukkan Nomor Handphone Customer : ")
		fmt.Scan(&customer.Phone)
		fmt.Print("Masukkan Alamat Customer : ")
		fmt.Scan(&customer.Address)
		fmt.Println("--------------------------------------------")
		err := validate.Struct(customer)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldErr := range validationErrors {
				if fieldErr.Field() == "Name" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "alpha":
						message = "Hanya boleh berupa huruf!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nama Customer", message)
				}
				if fieldErr.Field() == "Phone" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nomor Handphone Customer", message)
				}
				if fieldErr.Field() == "Address" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Alamat Customer", message)
				}
			}
			fmt.Println("--------------------------------------------")
			continue
		}
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Id Customer : %d\nNama Customer : %s\nNomor Hp : %s\nAlamat : %s\n", customer.Id, customer.Name, customer.Phone, customer.Address)
		for true {
			fmt.Print("Apakah Data Tersebut Sudah Benar? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
				entity.UpdateCustomer(customer)
				isLooping = false
				break
			} else if choice == "n" {
				isLooping = true
				break
			} else {
				fmt.Println("Pilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
}

/*
================================== DISPLAY DELETE CUSTOMER FUNCTION ==================================
-> UI Untuk Menghapus Data Customer
*/

func displayDeleteCustomer() {

	fmt.Println("Delete Data Customer")
	fmt.Println("===========================")
	customers := entity.GetAllCustomer()
	displayAllCustomer()
	fmt.Print("Masukkan Id Customer yang hendak di Hapus : ")
	fmt.Scan(&customer.Id)
	for _, cust := range customers {
		if customer.Id == cust.Id {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			entity.DeleteCustomer(customer.Id)
			break
		}
	}
}

/*
================================== DISPLAY ALL EMPLOYER FUNCTION ==================================
-> UI Untuk Menampilkan Semua Data Karyawan
*/

func displayAllEmployer() {
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
}

/*
================================== DISPLAY ADD EMPLOYER FUNCTION ==================================
-> UI Untuk Menambahkan Data Karyawan
*/

func displayAddEmployer() {
	var choice string
	isLooping := true
	fmt.Println("Tambah Karyawan Baru")
	fmt.Println("===========================")
	for isLooping {
		fmt.Print("Masukkan Nama Karyawan : ")
		fmt.Scan(&employer.Name)
		fmt.Print("Masukkan Nomor Handphone Karyawan : ")
		fmt.Scan(&employer.Phone)
		fmt.Print("Masukkan Alamat Karyawan : ")
		fmt.Scan(&employer.Address)
		fmt.Println("--------------------------------------------")
		err := validate.Struct(employer)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldErr := range validationErrors {
				if fieldErr.Field() == "Name" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "alpha":
						message = "Hanya boleh berupa huruf!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nama Karyawan", message)
				}
				if fieldErr.Field() == "Phone" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nomor Handphone Karyawan", message)
				}
				if fieldErr.Field() == "Address" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Alamat Karyawan", message)
				}
			}
			fmt.Println("--------------------------------------------")
			continue
		}
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Nama Karyawan : %s\nNomor Hp : %s\nAlamat : %s\n", employer.Name, employer.Phone, employer.Address)

		for true {
			fmt.Print("Apakah Data Tersebut Sudah Benar? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
				entity.AddEmployer(employer)
				isLooping = false
				break
			} else if choice == "n" {
				isLooping = true
				break
			} else {
				fmt.Println("Pilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
}

/*
================================== DISPLAY UPDATE EMPLOYER FUNCTION ==================================
-> UI Untuk Mengupdate Data Karyawan
*/

func displayUpdateEmployer() {
	var isLooping = true
	var choice string

	fmt.Println("Update Data Karyawan")
	fmt.Println("===========================")
	displayAllEmployer()
	fmt.Print("Masukkan Id Karyawan yang hendak di Update : ")
	fmt.Scan(&employer.Id)

	for isLooping {
		fmt.Print("Masukkan Nama Karyawan : ")
		fmt.Scan(&employer.Name)
		fmt.Print("Masukkan Nomor Handphone Karyawan : ")
		fmt.Scan(&employer.Phone)
		fmt.Print("Masukkan Alamat Karyawan : ")
		fmt.Scan(&employer.Address)
		fmt.Println("--------------------------------------------")
		err := validate.Struct(employer)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldErr := range validationErrors {
				if fieldErr.Field() == "Name" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "alpha":
						message = "Hanya boleh berupa huruf!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nama Karyawan", message)
				}
				if fieldErr.Field() == "Phone" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "min":
						message = "Harus mengandung minimal " + fieldErr.Param() + " karakter!"
					case "max":
						message = "Harus mengandung maksimal " + fieldErr.Param() + " karakter!"
					}
					fmt.Println("[!!!] Kolom Nomor Handphone Karyawan", message)
				}
				if fieldErr.Field() == "Address" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Alamat Karyawan", message)
				}
			}
			fmt.Println("--------------------------------------------")
			continue
		}
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Id Karyawan : %d\nNama Karyawan : %s\nNomor Hp : %s\nAlamat : %s\n", employer.Id, employer.Name, employer.Phone, employer.Address)
		for true {
			fmt.Print("Apakah Data Tersebut Sudah Benar? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
				entity.UpdateEmployer(employer)
				isLooping = false
				break
			} else if choice == "n" {
				isLooping = true
				break
			} else {
				fmt.Println("Pilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
}

/*
================================== DISPLAY DELETE EMPLOYER FUNCTION ==================================
-> UI Untuk Menghapus Data Karyawan
*/

func displayDeleteEmployer() {

	fmt.Println("Delete Data Employer")
	fmt.Println("===========================")
	employers := entity.GetAllEmployer()
	displayAllEmployer()
	fmt.Print("Masukkan Id Employer yang hendak di Hapus : ")
	fmt.Scan(&employer.Id)
	for _, emp := range employers {
		if employer.Id == emp.Id {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			entity.DeleteEmployer(employer.Id)
			break
		}
	}
}

/*
================================== DISPLAY ALL SERVICE FUNCTION ==================================
-> UI Untuk Menampilkan Semua Data Layanan
*/

func displayAllService() {
	services := entity.GetAllService()
	fmt.Println()
	fmt.Println("|\t\t\tLAYANAN ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Nama Layanan", "ID Layanan", "Harga", "Satuan"})
	for i, serv := range services {
		t.AppendRow([]interface{}{i + 1, serv.ServiceName, serv.Id, serv.Price, serv.Unit})
	}
	t.Render()

	fmt.Println()
}

/*
================================== DISPLAY ADD SERVICE FUNCTION ==================================
-> UI Untuk Menambahkan Data Layanan
*/

func displayAddService() {
	var choice string
	isLooping := true
	fmt.Println("Tambah Layanan Baru")
	fmt.Println("===========================")
	for isLooping {
		fmt.Print("Masukkan Nama Layanan : ")
		fmt.Scan(&layanan.ServiceName)
		fmt.Print("Masukkan Harga Layanan : ")
		fmt.Scan(&layanan.Price)
		fmt.Print("Masukkan Satuan : ")
		fmt.Scan(&layanan.Unit)
		fmt.Println("--------------------------------------------")
		err := validate.Struct(layanan)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldErr := range validationErrors {
				if fieldErr.Field() == "ServiceName" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Nama Layanan", message)
				}
				if fieldErr.Field() == "Price" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "numeric":
						message = "Harus Berupa Angka!"
					case "startsnotwith":
						message = "Tidak Boleh Diawali dengan " + fieldErr.Param() + "!"
					}
					fmt.Println("[!!!] Kolom Harga Layanan", message)
				}
				if fieldErr.Field() == "Unit" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Satuan", message)
				}
			}
			fmt.Println("--------------------------------------------")
			continue
		}
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Nama Layanan : %s\nHarga Layanan : %d\nSatuan : %s\n", layanan.ServiceName, layanan.Price, layanan.Unit)

		for true {
			fmt.Print("Apakah Data Tersebut Sudah Benar? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
				entity.AddService(layanan)
				isLooping = false
				break
			} else if choice == "n" {
				isLooping = true
				break
			} else {
				fmt.Println("Pilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
}

/*
================================== DISPLAY UPDATE SERVICE FUNCTION ==================================
-> UI Untuk Mengupdate Data Layanan
*/

func displayUpdateService() {
	var isLooping = true
	var choice string

	fmt.Println("Update Data Layanan")
	fmt.Println("===========================")
	displayAllService()
	fmt.Print("Masukkan Id Layanan yang hendak di Update : ")
	fmt.Scan(&layanan.Id)

	for isLooping {
		fmt.Print("Masukkan Nama Layanan : ")
		fmt.Scan(&layanan.ServiceName)
		fmt.Print("Masukkan Harga Layanan : ")
		fmt.Scan(&layanan.Price)
		fmt.Print("Masukkan Satuan : ")
		fmt.Scan(&layanan.Unit)
		fmt.Println("--------------------------------------------")
		err := validate.Struct(layanan)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldErr := range validationErrors {
				if fieldErr.Field() == "ServiceName" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Nama Layanan", message)
				}
				if fieldErr.Field() == "Price" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					case "numeric":
						message = "Harus Berupa Angka!"
					case "startsnotwith":
						message = "Tidak Boleh Diawali dengan " + fieldErr.Param() + "!"
					}
					fmt.Println("[!!!] Kolom Harga Layanan", message)
				}
				if fieldErr.Field() == "Unit" {
					var message string
					switch fieldErr.Tag() {
					case "required":
						message = "Tidak Boleh Kosong!"
					}
					fmt.Println("[!!!] Kolom Satuan", message)
				}
			}
			fmt.Println("--------------------------------------------")
			continue
		}
		fmt.Println("Data Yang Anda Inputkan =>")
		fmt.Printf("Id Layanan : %d\nNama Layanan : %s\nHarga Layanan : %d\nSatuan : %s\n", layanan.Id, layanan.ServiceName, layanan.Price, layanan.Unit)

		for true {
			fmt.Print("Apakah Data Tersebut Sudah Benar? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
				entity.UpdateService(layanan)
				isLooping = false
				break
			} else if choice == "n" {
				isLooping = true
				break
			} else {
				fmt.Println("Pilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
}

/*
================================== DISPLAY DELETE SERVICE FUNCTION ==================================
-> UI Untuk Menghapus Data Layanan
*/

func displayDeleteService() {

	fmt.Println("Delete Data Layanan")
	fmt.Println("===========================")
	services := entity.GetAllService()
	displayAllService()
	fmt.Print("Masukkan Id Layanan yang hendak di Hapus : ")
	fmt.Scan(&layanan.Id)
	for _, serv := range services {
		if layanan.Id == serv.Id {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
			entity.DeleteService(serv.Id)
			break
		}
	}
}

/*
================================== DISPLAY ALL TRANSACTION FUNCTION ==================================
-> UI Untuk Menampilkan Semua Data Transaksi
*/

func displayAllTransaction() {
	var choice string
	transactions := entity.GetAllTransaction()
	fmt.Println()
	fmt.Println("|\t\t\tTRANSAKSI ENIGMA LAUNDRY\t\t\t|")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	resetTable()
	t.AppendHeader(table.Row{"No", "Tanggal Masuk", "Tanggal Keluar", "ID Transaksi", "Nomor Transaksi", "Diterima Oleh", "Nama Customer", "No Hp Customer"})
	for i, trx := range transactions {
		dateIn, err := time.Parse("2006-01-02T15:04:05Z", trx.DateIn)
		dateOut, err := time.Parse("2006-01-02T15:04:05Z", trx.DateOut)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		trx.DateIn = dateIn.Format("2006-01-02")
		trx.DateOut = dateOut.Format("2006-01-02")
		t.AppendRow([]interface{}{i + 1, trx.DateIn, trx.DateOut, trx.Id, trx.NoTransaction, trx.EmployerName, trx.CustomerName, trx.CustomerPhone})
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

/*
================================== DISPLAY DETAIL TRANSACTION FUNCTION ==================================
-> UI Untuk Menampilkan Detail Data Dari Transaksi
*/

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

/*
================================== DISPLAY ADD TRANSACTION FUNCTION ==================================
-> UI Untuk Menambahkan Data Transaksi dan Detailnya
*/

func displayAddTransaction() {
	var choice string
	isLooping := true
	var arrays []entity.DetailTransaksi
	fmt.Println("Tambah Transaksi")
	fmt.Println("===========================")

	displayAllCustomer()
	displayAllEmployer()
	displayAllService()
	fmt.Println("FORM TRANSAKSI")
	fmt.Println("-----------------------------------------")
	transaksi.Id = entity.GetLastIdTransaction()
	fmt.Print("Masukkan Nomor Transaksi (000000) : ")
	fmt.Scan(&transaksi.NoTransaction)
	fmt.Print("Masukkan Tanggal Masuk (yyyy-mm-dd) : ")
	fmt.Scan(&transaksi.DateIn)
	fmt.Print("Masukkan Tanggal Keluar (yyyy-mm-dd) : ")
	fmt.Scan(&transaksi.DateOut)
	fmt.Print("Masukkan ID Karyawan : ")
	fmt.Scan(&transaksi.EmployerId)
	fmt.Print("Masukkan ID Customer : ")
	fmt.Scan(&transaksi.CustomerId)
	detail_transaksi.Id = entity.GetLastIdDetailTransaction()
	for isLooping {
		fmt.Println("\t-----------------------------------------")
		fmt.Println("\tFORM DETAIL TRANSAKSI")
		fmt.Println("\t-----------------------------------------")
		fmt.Print("\tMasukkan ID Layanan yang dipilih : ")
		fmt.Scan(&detail_transaksi.ServiceId)
		fmt.Print("\tJumlah : ")
		fmt.Scan(&detail_transaksi.Quantity)
		arrays = append(arrays, detail_transaksi)
		detail_transaksi.Id++
		for true {
			fmt.Print("\tIngin Menambah Detail Lainnya? (Y/n) ")
			fmt.Scan(&choice)
			if choice == "Y" {
				isLooping = true
				break
			} else if choice == "n" {
				isLooping = false
				break
			} else {
				fmt.Println("\tPilihan Tidak Valid. Harus Y atau n !!!")
			}
		}
	}
	entity.MakeTransaction(transaksi, arrays)
}

/*
================================== RESET TABLE FUNCTION ==================================
-> Untuk Mereset Header, Rows dan Footer pada Table yang digunakan dari Package go-pretty
*/

func resetTable() {
	t.ResetHeaders()
	t.ResetRows()
	t.ResetFooters()
}
