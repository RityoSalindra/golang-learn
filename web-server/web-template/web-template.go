package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Rityo Salindra",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("template.html") //Fungsi template.ParseFiles() digunakan untuk parsing template, mengembalikan 2 data yaitu instance template-nya dan error (jika ada).
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
