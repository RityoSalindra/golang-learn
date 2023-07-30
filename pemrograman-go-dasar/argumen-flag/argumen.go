package main

import (
	"fmt"
	"os"
)

func main() {
	var argsRaw = os.Args //Data arguments bisa didapat lewat variabel os.Args
	fmt.Printf("-> %#v\n", argsRaw)
	// -> []string{".../argumen", "banana", "potato", "ice cream"}

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
	// -> []string{"banana", "potatao", "ice cream"}
}
