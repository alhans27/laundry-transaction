-- INSERT DATA TO mst_customer
INSERT INTO mst_customer (id, name, phone, address) VALUES (1, 'Jessica', '081265498721', 'Jakarta');
INSERT INTO mst_customer (id, name, phone, address) VALUES (2, 'Mayang', '089467552211', 'Malang');
INSERT INTO mst_customer (id, name, phone, address) VALUES (3, 'Hanif', '085432998599', 'Surakarta');

SELECT * FROM mst_customer;

-- INSERT DATA TO mst_employer

INSERT INTO mst_employer (id, name, phone, address) VALUES (1, 'Mirna', '085891772952', 'Surabaya');
INSERT INTO mst_employer (id, name, phone, address) VALUES (2, 'Siti', '085891772445', 'Surabaya');
INSERT INTO mst_employer (id, name, phone, address) VALUES (3, 'Teguh', '085891772245', 'Surabaya');

SELECT * FROM mst_employer;

-- INSERT DATA TO mst_layanan

INSERT INTO mst_layanan (id, service_name, price, unit) VALUES (1, 'Cuci + Setrika', 7000, 'KG');
INSERT INTO mst_layanan (id, service_name, price, unit) VALUES (2, 'Laundry Bedcover', 50000, 'Buah');
INSERT INTO mst_layanan (id, service_name, price, unit) VALUES (3, 'Laundry Boneka', 25000, 'Buah');
INSERT INTO mst_layanan (id, service_name, price, unit) VALUES (4, 'Laundry Karpet', 60000, 'Buah');
INSERT INTO mst_layanan (id, service_name, price, unit) VALUES (5, 'Cuci', 5000, 'KG');
INSERT INTO mst_layanan (id, service_name, price, unit) VALUES (6, 'Laundry Jaket', 10000, 'KG');

SELECT * FROM mst_layanan;

-- INSERT DATA TO trx_bill

INSERT INTO trx_bill (id, no_trx, date_in, date_out, customer_id, employer_id) VALUES (1, '12345', '2022-08-18', '2022-08-20', 1, 1);
INSERT INTO trx_bill (id, no_trx, date_in, date_out, customer_id, employer_id) VALUES (2, '54321', '2022-08-18', '2022-08-20', 2, 1);

SELECT * FROM trx_bill;

-- INSERT DATA TO trx_bill_detail

INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES (1, 1, 5, 1);
INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES (2, 2, 1, 1);
INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES (3, 3, 2, 1);

INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES (4, 5, 4, 2);
INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES (5, 4, 3, 2);
INSERT INTO trx_bill_detail (id, service_id, quantity, trx_bill_id) VALUES (6, 6, 1, 2);

SELECT * FROM trx_bill_detail;