package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var gameOver bool
var word []string

func game() {
	var content string
	var expectation int
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
		check_win(n, select_word)
		board()
	}

	fmt.Println("\nВведите: \n1 - Чтобы начать новую игру \n2 - чтобы выйти из игры")
	for {
		fmt.Scan(&expectation)
		if expectation == 1 {
			main()
		} else if expectation == 2 {
			os.Exit(0)
		} else {
			fmt.Println("Вы ввели недопустимое значение. Попробуйте ещё раз:")
		}
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
