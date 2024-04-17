package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i < 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consuming %d\n", num)
		time.Sleep(1 * time.Second) // Simulate some work by the consumer
	}
}

func main() {
	normalChan := make(chan int)
	bufferedChan := make(chan int, 3)

	fmt.Println("Using normal channel")
	go producer(normalChan)
	consumer(normalChan)

	time.Sleep(2 * time.Second) //wait for the producer to finish

	fmt.Println("Using buffered channel:")
	go producer(bufferedChan)
	consumer(bufferedChan)

}
