package entity

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
	ServiceName string
	Price       int
	Unit        string
	Quantity    int
	TotalPrice  int
}
