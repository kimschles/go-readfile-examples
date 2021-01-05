package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Open the file
	file, err := os.Open("example.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create a reader with the encoding/csv package
	reader := csv.NewReader(file)
	for {
		// create a record for each comma delimited value
		record, err := reader.Read()
		// once you've reached the end of the file, break the loop
		if err == io.EOF {
			break
		}
		// If there was an error scanning the file, report it and exit
		if err != nil {
			log.Fatal(err)
		}
		// Print each record to the console
		fmt.Println(record)
	}

	// Close the file to free resources
	err = file.Close()
}
