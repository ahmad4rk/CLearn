
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Alat struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
	Status   string `json:"status"`
}

type Peminjaman struct {
	ID           int    `json:"id"`
	AlatID       int    `json:"alat_id"`
	NamaPeminjam string `json:"nama_peminjam"`
	Status       string `json:"status"`
}

var alat = []Alat{
	{1, "Laptop Asus", "Laptop", "tersedia"},
	{2, "Kamera Canon", "Kamera", "dipinjam"},
	{3, "Arduino Uno", "Elektronik", "tersedia"},
	{4, "Laptop Lenovo", "Laptop", "tersedia"},
	{5, "Kamera Sony", "Kamera", "tersedia"},
}

var peminjaman []Peminjaman

func main() {
	http.HandleFunc("/alat", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method tidak diizinkan", 405)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(alat)
	})

	http.HandleFunc("/peminjaman", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method tidak diizinkan", 405)
			return
		}

		var p Peminjaman
		if json.NewDecoder(r.Body).Decode(&p) != nil {
			http.Error(w, "Data tidak valid", 400)
			return
		}

		p.ID = len(peminjaman) + 1
		p.Status = "menunggu persetujuan"
		peminjaman = append(peminjaman, p)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(p)
	})

	fmt.Println("Server berjalan di localhost:8080")
	http.ListenAndServe(":8080", nil)
}
