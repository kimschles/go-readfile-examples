// Note: this code and comments are from Chapter 5 of [Head First Go by Jay McGavren](http://headfirstgo.com/)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the data file for reading
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	// create a new scanner for the file
	scanner := bufio.NewScanner(file)

	// loops until the end of the file is reached and scanner.Scan returns false
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// If there was an error scanning the file, report it and exit
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Close the file to free resources
	err = file.Close()

}
