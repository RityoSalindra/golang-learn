package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message)) //casting string ke byte
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world"
	w.Write([]byte(message)) //casting string ke byte
}

// Fungsi http.HandleFunc() digunakan untuk routing
func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())

	}

}
