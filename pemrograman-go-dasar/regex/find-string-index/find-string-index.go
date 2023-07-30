package main

import (
	"fmt"
	"regexp"
)

func main() {
	//mirip find string tapi yang dikembalikan indexnya
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]

	var str = text[0:6]
	fmt.Println(str)
	// "banana"
}
