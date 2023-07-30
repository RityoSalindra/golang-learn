package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

// bertugas melakukan request ke http://localhost:8080/users, menerima response dari request tersebut, lalu menampilkannya.
func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{} //&http.Client{} menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.
	var data []student

	request, err := http.NewRequest("GET", baseURL+"/users", nil)
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

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}
