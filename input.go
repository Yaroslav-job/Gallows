package main

import (
	"fmt"
	"unicode"
)

func input() string {
	var letter string
	fmt.Println("Введите букву:")
	for {
		fmt.Scan(&letter)
		if len(letter) == 1 && unicode.IsLetter(rune(letter[0])) {
			break
		} else {
			fmt.Println("Введите одну букву:")
		}
	}
	return letter
}
