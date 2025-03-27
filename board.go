package main

import (
	"fmt"
)

func word_start(select_word string) {
	len_word := len(select_word)
	word = make([]string, len_word)
	for i := 0; i < len_word; i++ {
		word[i] = "_"
	}
}

func board() {
	for i := 0; i < len(word); i++ {
		fmt.Printf("%s", word[i])
	}
	fmt.Printf("\n")
}
