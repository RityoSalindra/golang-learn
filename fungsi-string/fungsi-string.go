package main

import (
	"fmt"
	"strings"
)

func main() {
	//cek apakah parameter kedua adalah bagian dari parameter pertama
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	//deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	//deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua).
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	//untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama)
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	//untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4

	//digunakan untuk replace atau mengganti bagian dari string dengan string tertentu
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	//Digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"

	//Digunakan untuk memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua).
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	//Memiliki kegunaan berkebalikan dengan strings.Split().

	var data = []string{"banana", "papaya", "tomato"}
	var strng = strings.Join(data, "-")
	fmt.Println(strng) // "banana-papaya-tomato"

	//Mengubah huruf-huruf string menjadi huruf kecil.
	var str1 = strings.ToLower("aLAy")
	fmt.Println(str1) // "alay"

	//Mengubah huruf-huruf string menjadi huruf besar.
	var str2 = strings.ToUpper("eat!")
	fmt.Println(str2) // "EAT!"
}
