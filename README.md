ENIGMA LAUNDRY

# 1. Project Overview

Enigma Laundry merupakan aplikasi berbasis console yang digunakan untuk mencatat transaksi dari toko Enigma Laundry. Aplikasi ini memiliki 4 menu utama. Yaitu Menu Data Customer, Menu Data Karyawan, Menu Data Layanan dan Menu Data Transaksi. Aplikasi ini bisa melakukan VIEW, INSERT, UPDATE dan DELETE Data Customer, Karyawan dan Layanan. Selain itu Aplikasi ini bisa VIEW dan INSERT Data Transaksi.

# 2. Installation Guide

Aplikasi disertai file _DDL.sql_, _DML.sql_ dan ERDnya.
Sebelum menjalankan aplikasi, perhatikan hal-hal berikut: - Database yang digunakan adalah PostgresSQL - Pastikan Anda sudah membuat Database dengan nama "enigma-laundry" - Jalankan file DDL.sql pada Query Tool di Database untuk membuat Tabel yang dibutuhkan oleh aplikasi ini - Jalankan file DML.sql pada Query Tool di Database untuk membuat data dummy awal guna menjalankan aplikasi ini - Cari File bernama enigma-laundry.exe untuk menjalankan aplikasi, atau jika Anda adalah developer, bisa gunakan perintah _"go run main.go"_ di directory project ini

# Usage Guide

Aplikasi ini memiliki 4 Menu Utama yaitu Menu Data Customer, Menu Data Karyawan, Menu Data Layanan dan Menu Data Transaksi.
Masing-masing menu memiliki sub menu sebagai berikut:

Menu Data Customer

- Lihat Data Customer
- Tambah Data Customer
- Update Data Customer
- Delete Data Customer

Menu Data Karyawan

- Lihat Data Karyawan
- Tambah Data Karyawan
- Update Data Karyawan
- Delete Data Karyawan

Menu Data Layanan

- Lihat Data Layanan
- Tambah Data Layanan
- Update Data Layanan
- Delete Data Layanan

Menu Data Transaksi

- Lihat Semua Data Transaksi
  - Melihat Detail Tiap Transaksi
- Tambah Data Transaksi

1. Aplikasi pertama kali akan menampilkan Menu Utama
2. Untuk mengakses setiap menu, User diminta untuk memasukkan nomor dari menu yang diinginkan.
3. Terkadang user juga diminta untuk menginputkan (Y/n) ketika mengakses salah satu menu.

# MENU DATA CUSTOMER

# MENU TAMBAH DATA CUSTOMER

User akan diminta untuk menginput beberapa informasi untuk menambahkan Data Customer.

1. Nama Customer --> Nama Customer hanya berupa nama panggilan saja yang terdiri dari satu kata, bukan nama lengkap. Minimal berjumlah 3 karakter. Dan tidak boleh mengandung simbol maupun angka
2. Nomor HP Customer --> Nomor Hp Customer harus berupa dengan minimal karakter berjumlah 11 dan maksimal 13
3. Alamat Customer --> Alamat Customer dapat berupa huruf, angka maupun simbol. Namun hanya 1 kata saja
4. Akan ada konfirmasi terkait data yang dimasukkan oleh User apakah sudah sesuai atau belum, jika belum sesuai maka User bisa menginputkan _n_, jika sudah sesuai maka user bisa menginputkan _Y_
5. Apabila data berhasil ditambahkan, maka akan ada pesan _"Successfully Insert Data"_

# MENU UPDATE DATA CUSTOMER

Aplikasi pertama-tama akan menampilkan seluruh data Customer yang ada. Apabila User ingin mengubah salah satu data, maka:

1. User akan diminta untuk menginputkan _ID Customer_ yang hendak diubah datanya
2. Kemudian User diminta untuk menginputkan data Customer sesuai dengan ketentuan pada Tambah Data Customer
3. Apabila data berhasil diupdate, maka akan ada pesan _"Successfully Update Data"_

# DELETE DATA CUSTOMER

Aplikasi pertama-tama akan menampilkan seluruh data Customer yang ada. Apabila User ingin menghapus salah satu data, maka:

1. User akan diminta untuk menginputkan _ID Customer_ yang hendak dihapus
2. Apabila ID yang hendak dihapus tidak ada dalam Database, maka tidak akan terjadi apapun
3. Apabila Data Customer yang hendak dihapus telah tercatat melakukan transaksi, maka akan muncul pesan pemberitahuan dan data tersebut tidak dapat dihapus

# MENU DATA KARYAWAN

Menu ini memiliki ketentuan yang hampir sama dengan menu yang ada pada _MENU DATA CUSTOMER_

# MENU DATA LAYANAN

# MENU TAMBAH DATA LAYANAN

User akan diminta untuk menginput beberapa informasi untuk menambahkan Data Layanan.

1. Nama Layanan --> Nama Customer hanya terdiri dari satu kata. Minimal berjumlah 3 karakter.
2. Harga Layanan --> Harga Layanan harus berupa angka
3. Satuan --> Satuan merupakan cara sebuah layanan diberi harga, apakah per kilogram atau per buah
4. Akan ada konfirmasi terkait data yang dimasukkan oleh User apakah sudah sesuai atau belum, jika belum sesuai maka User bisa menginputkan _n_, jika sudah sesuai maka user bisa menginputkan _Y_
5. Apabila data berhasil ditambahkan, maka akan ada pesan _"Successfully Insert Data"_

# MENU UPDATE DATA LAYANAN

Aplikasi pertama-tama akan menampilkan seluruh data Layanan yang ada. Apabila User ingin mengubah salah satu data, maka:

1. User akan diminta untuk menginputkan _ID Layanan_ yang hendak diubah datanya
2. Kemudian User diminta untuk menginputkan data Layanan sesuai dengan ketentuan pada Tambah Data Layanan
3. Apabila data berhasil diupdate, maka akan ada pesan _"Successfully Update Data"_

# DELETE DATA LAYANAN

Aplikasi pertama-tama akan menampilkan seluruh data Layanan yang ada. Apabila User ingin menghapus salah satu data, maka:

1. User akan diminta untuk menginputkan _ID Layanan_ yang hendak dihapus
2. Apabila ID yang hendak dihapus tidak ada dalam Database, maka tidak akan terjadi apapun
3. Apabila Data Customer yang hendak dihapus telah tercatat melakukan transaksi, maka akan muncul pesan pemberitahuan dan data tersebut tidak dapat dihapus

# DATA TRANSAKSI

# LIHAT SEMUA DATA TRANSAKSI

1. Aplikasi pertama-tama akan menampilkan seluruh data Transaksi yang ada.
2. Akan muncul pilihan bagi User apakah ingin melihat detail transaksi atau tidak. User dapat memilih _"Y"_ apabila ingin melihat detail dari salah satu transaksi. Atau User dapat memilih _"n"_ apabila tidak ingin melihat detail transaksi
3. User yang ingin melihat Detail Transaksi akan diminta untuk memasukkan _ID Transaksi_ yang ingin dilihat.
4. Aplikasi akan menampilkan Detail Transaksi dari transaksi yang dipilih User. Apabila ID yang dimasukkan oleh User tidak ada dalam Database, maka aplikasi akan menampilkan detail kosong

# TAMBAH DATA TRANSAKSI

Aplikasi akan menampilkan Data Customer, Data Karyawan dan Data Layanan
User yang ingin menambahkan transaksi baru, harus menginputkan data yang mengacu pada tabel-tabel yang ditampilkan oleh aplikasi. Beberapa informasi yang harus diinputkan oleh User:

1. Nomor Transaksi --> Berupa Angka
2. Tanggal Masuk --> Tanggal Customer Melaundry-kan barangnya. Menggunakan format tahun-bulan-tanggal
3. Tanggal Keluar --> Tanggal rencana barang Customer selesai dikerjakan. Menggunakan format tahun-bulan-tanggal
4. ID Karyawan --> ID dari Karyawan yang menangani transaksi. Mengacu pada tabel yang ditampilkan oleh aplikasi
5. ID Customer --> ID dari Customer yang melakukan transaksi. Mengacu pada tabel yang ditampilkan oleh aplikasi

Setelah User menginputkan Informasi Umum terkait transaksi, User akan diminta memasukkan Detail Transaksi, berupa:

1. ID Layanan yang dipilih oleh Customer. ID merupakan ID layanan yang terdaftar dalam database dengan cara mengacu pada tabel yang ditampilkan oleh aplikasi
2. Jumlah barang yang dilaundrykan terkait layanan tersebut

# MENU EXIT

Untuk Keluar Dari Aplikasi

Copyright2024
Dev: Hanif Sajid Al Amna a.k.a Alhans
