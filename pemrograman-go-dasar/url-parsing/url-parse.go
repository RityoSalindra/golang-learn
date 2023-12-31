package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://kalipare.com:80/hello?name=Rityo Salindra&age=21"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)       // kalipare.com:80
	fmt.Printf("path: %s\n", u.Path)       // /hello

	var name = u.Query()["name"][0] // Rityo Salindra
	var age = u.Query()["age"][0]   // 21
	fmt.Printf("name: %s, age: %s\n", name, age)
}
