package main

import (
	"fmt"
	"strings"
)

var error_count int
var hangman = []string{
	`
|------------|
|            |
|            |
|            |
|            |
|            |
|            |
==============
`,
	`
|------------|
|            |
|         |  |
|         |  |
|         |  |
|         |  |
|         |  |
==============
`,
	`
|------------|
|    +====+  |
|         |  |
|         |  |
|         |  |
|         |  |
|         |  |
==============
`,
	`
|------------|
|    +====+  |
|    |    |  |
|    0    |  |
|         |  |
|         |  |
|         |  |
==============
`,
	`
|------------|
|    +====+  |
|    |    |  |
|    0    |  |
|    |    |  |
|         |  |
|         |  |
==============
`,
	`
|------------|
|    +====+  |
|    |    |  |
|    0    |  |
|   /|\   |  |
|         |  |
|         |  |
==============
`,
	`
|------------|
|    +====+  |
|    |    |  |
|    0    |  |
|   /|\   |  |
|   / \   |  |
|         |  |
==============
`,
}

func check_letter(select_word string, letter string) int {
	found := 1
	for i := 0; i < len(select_word); i++ {
		if string(select_word[i]) == letter {
			word[i] = letter
			found = 0
		}
	}
	return found
}

func check_win(n int, select_word string) {
	if n != 0 {
		error_count += 1
	}

	fmt.Printf("%s", hangman[error_count])

	if error_count == 6 {
		fmt.Println("Вы проиграли")
		gameOver = true
	}

	currentWord := strings.Join(word, "")
	if currentWord == select_word {
		fmt.Println("Вы победили")
		gameOver = true
	}
}
