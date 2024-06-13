package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var characters int

func main() {

	fileName := os.Args[1]

	contents, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	//number of characters
	characters = len(contents)

	fmt.Printf("characters: %d \n", characters)

	//number of words
	contentString := strings.Fields(string(contents))

	words := len(contentString)

	//number of lines 
	lines := strings.Split(string(contents),"\n")

	if len(os.Args) > 2 {
		for i := 0; i< len(os.Args); i++ {
			if os.Args[i] == "-w" {
				fmt.Printf("Words: %d \n",words)
			}  
			if os.Args[i] == "-l" {
				fmt.Printf("Lines: %d \n", len(lines))
			}
		}
	} else {
		return
	}

}