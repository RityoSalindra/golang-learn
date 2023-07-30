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

	var data User

	var err = json.Unmarshal(jsonData, &data) //dengan json string tersebut dimasukan ke statement fungsi tersebut.
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)
}
