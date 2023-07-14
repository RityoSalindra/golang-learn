package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// t1 := time.Now()
	time.Sleep(5 * time.Second)
	// t2 := time.Now()

	duration := time.Since(start) //waktu yang digunakan untuk eksekusi setelah start

	// duration := t2.Sub(t1)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())

}
