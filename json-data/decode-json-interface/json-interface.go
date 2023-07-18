package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"` //untuk membuat variabel baru penampung hasil decode json string.
	Age      int
}

func main() {
	var jsonString = `{"Name": "Rityo Salindra", "Age": 21}`
	var jsonData = []byte(jsonString)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{}) //Variabel bertipe interface{} juga bisa digunakan untuk menampung hasil decode.
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])
}
