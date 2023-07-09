package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "strawberry"
	baFruits[0] = "strawberry"

	fmt.Println(fruits)   // [apple strawberry banana melon]
	fmt.Println(aFruits)  // [apple strawberry banana]
	fmt.Println(bFruits)  // [strawberry banana melon]
	fmt.Println(aaFruits) // [strawberry]
	fmt.Println(baFruits) // [strawberry]

}
