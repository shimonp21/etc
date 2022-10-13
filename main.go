package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print((time.Second * 50) / (time.Second * 70))
}

func reader(channel <-chan int) {
	time.Sleep(10 * time.Second)

	for num := range channel {
		fmt.Println("reader:", num)
		time.Sleep(1 * time.Second)
	}
}
