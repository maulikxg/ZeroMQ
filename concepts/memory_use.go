package main

import (
	"fmt"
	"os"
)

func main() {

	const fileSize = 17 * 1024 * 1024 * 1024

	data := make([]byte, fileSize)

    // open the file 
	file , err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	// write the data to the file 
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}

	fmt.Println("File written successfully")
}