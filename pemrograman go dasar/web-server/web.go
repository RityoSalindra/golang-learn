package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index) //Fungsi http.HandleFunc() digunakan untuk routing aplikasi web

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil) //digunakan untuk menghidupkan server sekaligus menjalankan aplikasi menggunakan server tersebut. Di Go
}
