package main

import (
	"fmt"
)

func main() {
	clear()
	gameOver = false
	error_count = 0
	fmt.Println("Добро пожаловать в игру \"Висилица\"")
	game()
}
