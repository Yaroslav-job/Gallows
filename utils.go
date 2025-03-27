package main

import (
	"math/rand"
	"os"
	"os/exec"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func random(n int) int {
	return rand.Intn(n)
}
