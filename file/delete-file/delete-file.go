package main

import (
	"fmt"
	"os"
)

var path = "D:/DATA/agree/learn-golang/file/delete-file/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di delete")
}

func main() {
	deleteFile()
}
