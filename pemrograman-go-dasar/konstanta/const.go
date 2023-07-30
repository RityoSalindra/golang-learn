package main

import "fmt"

func main() {
	//konstanta yaitu variabel yang tidak dapat diubah
	const firstName string = "john"

	//multi konstanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	fmt.Print("halo ", firstName, "!\n")

}
