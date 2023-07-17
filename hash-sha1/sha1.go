package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var text = "this is secret"
	var sha = sha1.New()
	sha.Write([]byte(text))                            //digunakan untuk menge-set data yang akan di-hash. Data harus dalam bentuk []byte
	var encrypted = sha.Sum(nil)                       //digunakan untuk eksekusi proses hash, menghasilkan data yang sudah di-hash dalam bentuk []byte , membutuhkan parameter ,isi dengan nil
	var encryptedString = fmt.Sprintf("%x", encrypted) //bentuk heksadesimal string dari data yang sudah di-hash, bisa memanfaatkan fungsi fmt.Sprintf dengan layout format %x

	fmt.Println(encryptedString)
	// f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8
}
