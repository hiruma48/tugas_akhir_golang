package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type daftar_siswa struct {
	ID    string
	Nama  string
	Umur  int
	Nilai int
}

func koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/db_belajar_sql")

	if err != nil {
		return nil, err
	}
	return db, nil
}

var data []daftar_siswa

func main() {
	ambil_data()
	http.HandleFunc("/siswa", data_siswa)
	http.HandleFunc("/Cari_siswa", cari_siswa)

	fmt.Println("menjalankan web server pada localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func data_siswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)

}
func cari_siswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var namasiswa = r.FormValue("Nama")
		var result []byte
		var err error

		for _, each := range data {
			if each.Nama == namasiswa {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}

		}
		http.Error(w, "data siswa tidak tersedia", http.StatusBadRequest)
		return

	}
	http.Error(w, "", http.StatusBadRequest)
}
func ambil_data() {
	db, err := koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select *from tbl_pelajar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = daftar_siswa{}
		var err = rows.Scan(&each.ID, &each.Nama, &each.Umur, &each.Nilai)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
