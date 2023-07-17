package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "https://kalipare.com/"

	//Khusus encode data string yang isinya merupakan URL, lebih efektif menggunakan URLEncoding dibandingkan StdEncoding.

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
