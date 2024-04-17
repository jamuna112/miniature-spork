package main

/*
Buffered Channel:
A buffered channel is a type of channel in Go that has a fixed capacity
to hold a certain number of elements. When you create a buffered channel,
 you specify the capacity as the second argument to the make function.
 For example, ch := make(chan int, 10). In this case, the channel ch can
 hold up to 10 elements.
*/

// func main() {
// 	//create a buffered channel with a capacity of 3
// 	ch := make(chan int, 3)

// 	ch <- 1 // send data to the channel
// 	ch <- 2 //send more data
// 	ch <- 3 //send even more data

// 	//ch <- 4 //sending a 4th value would cause a deadlock because the channel is full

// 	fmt.Println(<-ch) //receive data from the channel
// 	fmt.Println(<-ch) // receive more data
// 	fmt.Println(<-ch) // receive even more data
// 	//fmt.Println(<-ch) // receive a 4th value would cause a deadlock because the channel is empty

// }
