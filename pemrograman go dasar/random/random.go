package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// randomizer := rand.New(rand.NewSource(10))    //membuat object randomizer sekaligus penentuan nilai seed-nya
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano())) //unique seed agar angka random yang dihasilkan berbeda
	fmt.Println("random ke-1:", randomizer.Int())
	fmt.Println("random float32:", randomizer.Float32()) //randomizer dengan tipe data yang berbeda
	fmt.Println("random uint:", randomizer.Uint32())
}
