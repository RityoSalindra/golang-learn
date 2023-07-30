package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(object) //digunakan untuk encoding data ke json string
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData) //hasil berupa []byte , sehingga di casting dulu menjadi string
	fmt.Println(jsonString)
}
