Deskripsi File
1. main.go
File main.go berisi implementasi fungsi untuk mengelompokkan kata-kata anagram dan sebuah program utama untuk menguji fungsi tersebut. Fungsi Group digunakan untuk mengelompokkan array string menjadi kumpulan anagram, sedangkan fungsi filter digunakan untuk mengurutkan karakter dalam sebuah string. Program utama menggunakan fungsi Group untuk mengelompokkan kata-kata anagram dari sebuah array yang telah ditentukan.

2. sql_query.sql
File sql_query.sql berisi sebuah query SQL untuk melakukan operasi join antara dua tabel. Query ini mengambil data dari tabel testt dan melakukan LEFT JOIN dengan tabel testt itu sendiri menggunakan kolom id_parent sebagai kunci join. Hasil query menampilkan id_name, name, dan parent_name dari tabel.

Cara Penggunaan
Untuk menggunakan file main.go, Anda dapat menjalankan program tersebut menggunakan perintah go run main.go. Pastikan untuk memberikan input array yang sesuai jika ingin menguji fungsi Group.

Untuk menggunakan file sql_query.sql, jalankan query tersebut di database Anda setelah mengganti nama tabel dengan nama tabel yang benar di dalam database.
