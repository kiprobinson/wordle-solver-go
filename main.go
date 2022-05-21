package main

import (
	"fmt"
	"wordle-solver-go/pkg/wordlist"
)

func main() {
	fmt.Println("wordlist:")
	fmt.Println(wordlist.GetWordList())
}
