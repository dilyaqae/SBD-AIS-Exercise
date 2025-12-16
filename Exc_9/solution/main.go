package main

// imports
import (
	"exc9/mapred"
	"fmt"
	"os"
	"strings"
)

// Main function
func main() {
	// todo read file
	data, err := os.ReadFile("res/meditations.txt")
	if err != nil {
		panic(err)
	}
	// convert into []string
	text := strings.Split(string(data), "\n")
	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	// todo print your result to stdout
	for word, count := range results {
		fmt.Printf("%s: %d\n", word, count)
	}
}
