package main

import (
	"fmt"
	"io"
	"os"
)

var path = "D:/DATA/agree/learn-golang/file/read-file/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644) //os.O_RDONLY maksudnya adalah read only.
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func main() {
	readFile()
}
