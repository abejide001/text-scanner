package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:] // read the input from the command line
	if len(args) <= 0 {
		fmt.Println("enter a value")
		return
	}

	in := bufio.NewScanner(os.Stdin) // read the input
	in.Split(bufio.ScanWords)        // scan for the words

	words := make(map[string]bool) // the key represents the unique word
	for in.Scan() {
		word := strings.ToLower(in.Text()) // convert the word to lowercase

		if len(word) > 2 { // check if the length of the word is greater than 2
			words[word] = true // specify the value to be true
		}
	}
	fmt.Println(words)

	query := args[0]
	if words[query] { // check if a given word exist
		fmt.Println("contains the input", query)
		return
	}
	fmt.Println("Does not contain the input", query)
}
