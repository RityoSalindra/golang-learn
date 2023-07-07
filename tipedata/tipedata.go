package main

import "fmt"

func main() {

	// numerik non desimal

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//numerik desimal

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//bool
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	//string

	var message = `Nama saya "Rityo".
				Salam kenal.
				Mari belajar "Golang".`

	fmt.Println(message)

}
