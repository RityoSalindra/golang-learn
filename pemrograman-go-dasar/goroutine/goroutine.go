package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2) //menentukan jumlah core saat eksekusi program

	go print(5, "halo") //penambahan go menandakan pembuatan gorountine baru
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input) //Scanln mengcapture karakter dan disimpan ke dalam variabel
}
