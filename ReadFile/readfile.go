package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args is a slice of strings representing the command-line arguments
	//This line excludes the program name
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {
		//This line attempts to open the said file
		//the line below returns a file descriptor "file" and error if there is an error.
		file, err := os.Open("ques8.txt")
		//If error is there
		if err != nil {
			fmt.Println(err.Error())
		}
		//An array is created of size 21
		arr := make([]byte, 21)

		//This line reads up to 21 bytes from file into the slice
		file.Read(arr)
		//Here the slice is converted into a string and is printed.
		fmt.Println(string(arr))
		//The file is closed
		file.Close()
	}

}
