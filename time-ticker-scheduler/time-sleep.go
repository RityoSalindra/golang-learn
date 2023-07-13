package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")

	for true {
		fmt.Println("Hello !!") //scheduler untuk menampilkan pesan halo setiap 1 detik
		time.Sleep(1 * time.Second)
	}

	var timer = time.NewTimer(4 * time.Second) // mengembalikan objek bertipe *time.Timer yang memiliki property C yang bertipe channel.
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

}
