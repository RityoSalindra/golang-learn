package main

import "fmt"

func main() {
	var firstName string = "Rityo"

	lastName := "Salindra" //deklarasi variabel inference(tanpa tipe data) biasa digunakan untuk multivariable dengan tipe data berbeda

	fmt.Printf("halo %s %s!\n", firstName, lastName) //printf tidak menghasilkan line baru jadi perlu \n dan outputnya didefinisikan diawal

	fmt.Println("halo", firstName, lastName+"!")

}
