package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create file
	newFile, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := newFile.Write([]byte("hello world and create files using Golang\n"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File creating sucessfully! Your size is %d bytes\n", size)
	newFile.Close()

	// Read file
	myFile, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Detected text on my file:", string(myFile))

	// Read file Specific bytes if my memmory limited
	myFileToRead, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(myFileToRead)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))

	}

	// Remove files
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}