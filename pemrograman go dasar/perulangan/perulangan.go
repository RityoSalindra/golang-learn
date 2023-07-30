package main

import "fmt"

func main() {

	//perulangan for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//perulangan for dengan argumen

	var x = 0

	for x < 10 {
		fmt.Println("Angka", x)
		x++
	}

	// for tanpa argumen
	var y = 0
	for {
		fmt.Println("Angka", y)
		y++

		if y == 10 {
			break
		}
	}

	//

}
