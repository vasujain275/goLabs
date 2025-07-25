package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a sentence: ")

	if !scanner.Scan() {
		fmt.Println("Failed to read input.")
		return
	}

	sentence := scanner.Text()
	words := strings.Fields(sentence) // handles multiple spaces, trims whitespace

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	// Print sorted by word for consistent output
	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s - %d\n", k, freq[k])
	}
}
