package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
the buffer belongs to the byte package of the Go language, and we can use these packages to manipulate the byte of the string.
For example, suppose we have a string. We can read the length of the string with the len function, which will return the numeric
length, but what if the strings are too large? We want to calculate in the form of chunks of data,
so in such situations, we can use the buffer; the buffer allows us to handle any size of the dynamic length, making it flexible.
*/

// example 1
func bufferExam1() {
	//creating the bytes package so that buffer can be used
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Testing ")
	strBuffer.WriteString("Test ")
	strBuffer.WriteString("For The Test..")

	fmt.Println("The string buffer output is ", strBuffer.String())
}

func bufferExam2() {

	//Creating buffer variable to hold and manage the string data
	var strBuffer bytes.Buffer
	fmt.Fprintf(&strBuffer, "It is number: %d, This is a string %v\n", 10, "Bridge")
	strBuffer.WriteString("[DONE]")
	fmt.Println("The string buffer output is", strBuffer.String())

}

func bufferExam3() {
	// creating buffer variable to hold and manage the string data
	var byteString bytes.Buffer
	byteString.Write([]byte("Hello "))
	fmt.Fprintf(&byteString, "Hello friends how are you?")
	byteString.WriteTo(os.Stdout)
}

func bufferExam4() {
	//creating buffer variable to hold and manage the string data
	var strByte bytes.Buffer
	strByte.Grow(64)
	strBytesVar := strByte.Bytes()
	strByte.Write([]byte("it is a 64 byte"))
	fmt.Printf("\n%b", strBytesVar[:strByte.Len()])
}

func bufferExam5() {
	var stryByte bytes.Buffer
	stryByte.Grow(64)
	stryByte.Write([]byte("kumar"))
	fmt.Printf("\n%b", stryByte.Len())
}

func bufferExam6() {
	var strByyte = []byte("Ranjan, Kumar")
	strByyte = bytes.TrimPrefix(strByyte, []byte("Hello, "))
	strByyte = bytes.TrimPrefix(strByyte, []byte("Good we will meet again,"))
	fmt.Printf("\nHello %s", strByyte)
}

func main() {
	bufferExam1()
	bufferExam2()
	bufferExam3()
	bufferExam4()
	bufferExam5()
	bufferExam6()
}
