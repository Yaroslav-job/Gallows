package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var gameOver = false
var word []string

func game() {
	var content string
	file, err := os.Open("dictionary.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}

	list := strings.Split(content, ", ")
	select_word := choosing_word(list)
	word_start(select_word)
	board()

	for !gameOver {
		letter := input()
		n := check_letter(select_word, letter)
		clear()
		board()
		check_win(n, select_word)
	}
}

func choosing_word(list []string) string {
	var word string
	n := len(list)
	if n == 0 {
		fmt.Println("Список слов пуст")
	}
	rand_word := random(n)
	word = list[rand_word]

	return word
}
