package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Rityo", "Salindra"}
	fungsi("halo", names)
	fungsi("dan", names)
}

func fungsi(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
