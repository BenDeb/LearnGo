package main

import (
	//"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	//"io"
	//"log"
	"os"
	"strings"
)

func main() {

	var filenamePtr = flag.String("filename", "problems.csv", "the name of the CSV file to load")
	flag.Parse()
	//Load the CSV file with os.Open()
	csvFile, err := os.Open(*filenamePtr)
	if err != nil {
		//log.Fatal(err)
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *filenamePtr))
	}
	//Close the file at the end of the function
	defer csvFile.Close()

	//Create a new reader
	csvr := csv.NewReader(csvFile)
	//problemNb := 0
	//new reader for stdin input
	//reader := bufio.NewReader(os.Stdin)

	//Cette première méthode ouvre une ligne du CSV à la fois
	/*for {
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
	}*/
	//Cette deuxième méthode ouvre tout le fichier en mémoire
	rows, err := csvr.ReadAll()
	if err != nil {
		exit("Failzd to parse the provided CSV file")
	}

	problems := parseRows(rows)
	score := checkAnswer(problems)
	fmt.Println("Final score:", score, "out of", len(problems))
}

func checkAnswer(problems []problem) int {
	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem number %d: %s\n", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.a {
			score++
		}
	}
	return score
}

func parseRows(rows [][]string) []problem {
	ret := make([]problem, len(rows))
	for i, row := range rows {
		ret[i] = problem{
			q: row[0],
			a: strings.TrimSpace(row[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
