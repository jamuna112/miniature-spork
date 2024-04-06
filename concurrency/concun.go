package main

import (
	"fmt"
	"time"
)

func foo(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, ":", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go foo("Jamuna")
	foo("Subedi")
}
