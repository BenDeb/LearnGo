package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	var filenamePtr = flag.String("filename", "problems.csv", "the name of the CSV file to load")
	flag.Parse()
	//Load the CSV file with os.Open()
	csvFile, err := os.Open(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}
	//Close the file at the end of the function
	defer csvFile.Close()

	//Create a new reader
	csvr := csv.NewReader(csvFile)
	score := 0
	problemNb := 0
	//new reader for stdin input
	reader := bufio.NewReader(os.Stdin)

	for {
		row, err := csvr.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("Problem %d : %v = %v\n", problemNb, row[0], row[1])
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.Compare(text, row[1]) == 0 {
			score += 1
		}

		problemNb++
	}
	fmt.Println("Final score:", score, "out of ", problemNb)
}
