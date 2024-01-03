CREATE TABLE mst_customer (
	id INTEGER NOT NULL,
	name VARCHAR(50),
	phone VARCHAR(15),
	address VARCHAR(100),
	PRIMARY KEY(id)
);

CREATE TABLE mst_employer (
	id INTEGER NOT NULL,
	name VARCHAR(50),
	phone VARCHAR(15),
	address VARCHAR(100),
	PRIMARY KEY(id)
);

CREATE TABLE mst_layanan (
	id INTEGER NOT NULL,
	service_name VARCHAR(100),
	price INTEGER,
	unit VARCHAR(50),
	PRIMARY KEY(id)
);

CREATE TABLE trx_bill (
	id INTEGER NOT NULL,
	no_trx VARCHAR(100),
	date_in DATE,
	date_out DATE,
	customer_id INTEGER NOT NULL,
	employer_id INTEGER NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(customer_id) REFERENCES mst_customer(id),
	FOREIGN KEY(employer_id) REFERENCES mst_employer(id)
);

CREATE TABLE trx_bill_detail (
	id INTEGER NOT NULL,
	service_id INTEGER NOT NULL,
	quantity INTEGER,
	trx_bill_id INTEGER,
	PRIMARY KEY(id),
	FOREIGN KEY(service_id) REFERENCES mst_layanan(id),
	FOREIGN KEY(trx_bill_id) REFERENCES trx_bill(id)
);