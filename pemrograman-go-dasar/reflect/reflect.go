package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe  variabel :", reflectValue.Type())
	fmt.Println("nilai variabel :", reflectValue.Interface()) //mengembalikan nilai interface kosong

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	var nilai = reflectValue.Interface().(int)
	fmt.Println("nilai asli (int) :", nilai)
}
