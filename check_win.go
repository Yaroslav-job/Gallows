package main

import (
	"fmt"
	"strings"
)

var error_count int

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
		switch error_count {
		case 1:
			fmt.Println("1 ошибка из 6")
		case 2:
			fmt.Println("2 ошибки из 6")
		case 3:
			fmt.Println("3 ошибки из 6")
		case 4:
			fmt.Println("4 ошибки из 6")
		case 5:
			fmt.Println("5 ошибок из 6")
		case 6:
			fmt.Println("6 ошибок из 6. Игра окончена")
			gameOver = true
		}
	}

	currentWord := strings.Join(word, "")
	if currentWord == select_word {
		fmt.Println("Вы победили")
		gameOver = true
	}
}
