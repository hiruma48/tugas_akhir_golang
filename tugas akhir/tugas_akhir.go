package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type daftar_siswa struct {
	ID    string
	Nama  string
	Umur  int
	Nilai int
}

func ambil_api() ([]daftar_siswa, error) {
	var err error
	var client = &http.Client{}
	var data []daftar_siswa

	request, err := http.NewRequest("POST", baseURL+"/siswa", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func cari_siswa(nama string) (daftar_siswa, error) {
	var err error
	var client = &http.Client{}
	var data daftar_siswa

	var param = url.Values{}
	param.Set("Nama", nama)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/Cari_siswa", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	// var siswa, err = ambil_api()
	// if err != nil {
	// 	fmt.Println("nama siswa tidak ada", err.Error())
	// 	return
	// }

	// for _, each := range siswa {
	// 	fmt.Printf("ID: %s\t Nama: %s\t Umur: %d\n Nilai: %d\n", each.ID, each.Nama, each.Umur, each.Nilai)
	// }

	var nama, err = cari_siswa("shani")
	if err != nil {
		fmt.Println("nama siswa tidak ada", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Nama: %s\t Umur: %d\n Nilai: %d\n", nama.ID, nama.Nama, nama.Umur, nama.Nilai)

}
